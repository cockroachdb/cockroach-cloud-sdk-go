#!/usr/bin/env bash
set -euo pipefail

# Sync OpenAPI spec from managed-service PR and regenerate client.
#
# This script orchestrates the entire OpenAPI sync workflow:
# 1. Validates event payload
# 2. Resolves managed-service PR metadata
# 3. Finds or creates SDK branch
# 4. Syncs OpenAPI spec and regenerates client
# 5. Updates changelog with Claude
# 6. Commits and pushes changes
# 7. Creates or updates SDK PR
#
# Usage:
#   sync-openapi.sh
#
# Required environment variables:
#   EVENT_TYPE: "openapi-spec-changed" or "openapi-spec-merged"
#   MANAGED_SERVICE_PR_URL: URL of the managed-service PR
#   MANAGED_SERVICE_SHA: Commit SHA (required for openapi-spec-merged)
#   MANAGED_SERVICE_TOKEN: GitHub token with managed-service read access
#   FORK_PUSH_TOKEN: GitHub token with fork push access
#   FORK_OWNER: GitHub username that owns the fork
#   CREATE_PR_TOKEN: GitHub token for creating PRs
#   GITHUB_REPOSITORY: Repository in owner/repo format

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/logging.sh"
source "$SCRIPT_DIR/lib/validation.sh"
source "$SCRIPT_DIR/lib/find-sdk-pr.sh"

# Validate required environment variables
validate_env() {
  local required_vars=(
    "EVENT_TYPE"
    "MANAGED_SERVICE_PR_URL"
    "MANAGED_SERVICE_TOKEN"
    "FORK_PUSH_TOKEN"
    "FORK_OWNER"
    "CREATE_PR_TOKEN"
    "GITHUB_REPOSITORY"
  )

  for var in "${required_vars[@]}"; do
    if [[ -z "${!var:-}" ]]; then
      log_error "Required environment variable $var is not set"
      exit 1
    fi
  done

  if [[ "$EVENT_TYPE" == "openapi-spec-merged" && -z "${MANAGED_SERVICE_SHA:-}" ]]; then
    log_error "MANAGED_SERVICE_SHA is required for openapi-spec-merged event"
    exit 1
  fi
}

# Resolve managed-service PR metadata from URL
resolve_managed_service_pr() {
  if [[ ! "$MANAGED_SERVICE_PR_URL" =~ github\.com/([^/]+)/([^/]+)/pull/([0-9]+) ]]; then
    log_error "Invalid PR URL format: $MANAGED_SERVICE_PR_URL"
    exit 1
  fi

  MS_OWNER="${BASH_REMATCH[1]}"
  MS_REPO="${BASH_REMATCH[2]}"
  MS_PR_NUMBER="${BASH_REMATCH[3]}"

  log_info "Managed-service PR: $MS_OWNER/$MS_REPO#$MS_PR_NUMBER"

  # Export for use in other functions
  export MS_OWNER MS_REPO MS_PR_NUMBER
}

# Configure Git
configure_git() {
  git config user.name "crl-console-bot"
  git config user.email "crl-console-bot@users.noreply.github.com"

  FORK_URL="https://github.com/$FORK_OWNER/$(basename "$GITHUB_REPOSITORY").git"
  log_info "Adding fork remote: $FORK_URL"
  git remote add fork "$FORK_URL"
}

# Determine or create SDK branch
# Two scenarios:
# 1. openapi-spec-changed: Create new SDK PR or update existing one
# 2. openapi-spec-merged: Update existing SDK PR (must already exist)
determine_sdk_branch() {
  export GIT_ASKPASS="$SCRIPT_DIR/lib/git-askpass.sh"
  export GIT_FORK_USER="$FORK_OWNER"
  export GIT_FORK_PASSWORD="$FORK_PUSH_TOKEN"
  export GIT_TERMINAL_PROMPT=0

  chmod +x "$GIT_ASKPASS"

  # Search all open PRs for one with a commit containing the managed-service PR URL
  if find_sdk_pr "$MANAGED_SERVICE_PR_URL"; then
    SDK_BRANCH="$SDK_PR_BRANCH"
    log_info "Using existing SDK PR branch: $SDK_BRANCH"

    git fetch fork "$SDK_BRANCH" || {
      log_error "Could not fetch branch $SDK_BRANCH from fork"
      exit 1
    }
    git checkout -B "$SDK_BRANCH" "fork/$SDK_BRANCH"
    EXISTING_SDK_PR="true"
  else
    # No existing SDK PR found
    if [[ "$EVENT_TYPE" == "openapi-spec-merged" ]]; then
      # For merged events, we expect an SDK PR to already exist from the earlier
      # openapi-spec-changed event. If not found, something went wrong.
      log_error "No SDK PR found for openapi-spec-merged event"
      exit 1
    fi

    # For changed events, create a new branch for a new SDK PR
    TIMESTAMP=$(date --utc +%Y%m%d-%H%M%S)
    SDK_BRANCH="automation/openapi-spec-change-$TIMESTAMP"
    log_info "Creating new SDK branch: $SDK_BRANCH"
    git checkout -b "$SDK_BRANCH"
    EXISTING_SDK_PR="false"
  fi

  export SDK_BRANCH EXISTING_SDK_PR
}

# Sync OpenAPI spec and regenerate client
# Fetches the spec from managed-service and regenerates SDK client code.
# For openapi-spec-changed events: fetch from PR branch
# For openapi-spec-merged events: fetch from the exact merged commit SHA
sync_and_regenerate() {
  SPEC_PATH_IN_MS="console/ui/cc-console/assets/docs/api/latest/openapi.json"
  LOCAL_SPEC_PATH="internal/spec/openapi.json"

  # Use commit SHA if provided (merged event), otherwise use PR head (changed event)
  if [[ -n "${MANAGED_SERVICE_SHA:-}" ]]; then
    REF="$MANAGED_SERVICE_SHA"
    log_info "Fetching openapi.json from $MS_OWNER/$MS_REPO at commit $MANAGED_SERVICE_SHA"
  else
    REF="refs/pull/$MS_PR_NUMBER/head"
    log_info "Fetching openapi.json from $MS_OWNER/$MS_REPO PR #$MS_PR_NUMBER"
  fi

  # Fetch the OpenAPI spec file from managed-service
  export GH_TOKEN="$MANAGED_SERVICE_TOKEN"
  gh api "repos/$MS_OWNER/$MS_REPO/contents/$SPEC_PATH_IN_MS?ref=$REF" \
    --jq '.content' | base64 --decode > "$LOCAL_SPEC_PATH.new"

  # Compare with current spec to see if anything changed
  if cmp --silent "$LOCAL_SPEC_PATH" "$LOCAL_SPEC_PATH.new"; then
    log_info "No changes to $LOCAL_SPEC_PATH"
    rm "$LOCAL_SPEC_PATH.new"
    HAS_CHANGES="false"
  else
    # Spec changed - replace it and regenerate the client code
    log_info "Spec file changed, replacing $LOCAL_SPEC_PATH"
    mv "$LOCAL_SPEC_PATH.new" "$LOCAL_SPEC_PATH"

    # Regenerate SDK client from the new spec using Docker
    # Note: sudo is required because docker containers run as root by default,
    # creating files owned by root. We then chown them back to the current user.
    log_info "Running make generate-openapi-client"
    sudo make generate-openapi-client
    sudo chown --recursive "$(id --user):$(id --group)" internal/openapi-generator/ pkg/client/ docs/ README.md

    # Check if regeneration produced any changes
    if [[ -z "$(git status --porcelain)" ]]; then
      HAS_CHANGES="false"
    else
      HAS_CHANGES="true"
    fi
  fi

  export HAS_CHANGES
}

# Update changelog with Claude
# Uses Claude AI to analyze the git diff, update CHANGELOG.md, and generate
# commit/PR metadata (title, body, description).
update_changelog() {
  if [[ "$HAS_CHANGES" != "true" ]]; then
    return
  fi

  log_info "Setting up Claude CLI..."
  curl --fail --silent --show-error --location https://claude.ai/install.sh | bash -s -- "latest"
  log_info "Claude CLI installed: $(claude --version)"

  # Configure Claude to use Vertex AI for authentication
  log_info "Invoking Claude to update changelog..."
  export CLAUDE_CODE_USE_VERTEX="1"
  export ANTHROPIC_VERTEX_PROJECT_ID="vertex-model-runners"
  export CLOUD_ML_REGION="us-east5"

  # Call helper script that invokes Claude and returns structured output
  CLAUDE_OUTPUT=$("$SCRIPT_DIR/lib/invoke-claude-changelog.sh" "$MANAGED_SERVICE_PR_URL" "${MANAGED_SERVICE_SHA:-}")

  # Parse the structured output (delimited by EOF markers)
  COMMIT_TITLE=$(echo "$CLAUDE_OUTPUT" | awk '/commit_title<<EOF_COMMIT_TITLE/{flag=1;next}/EOF_COMMIT_TITLE/{flag=0}flag')
  COMMIT_BODY=$(echo "$CLAUDE_OUTPUT" | awk '/commit_body<<EOF_COMMIT_BODY/{flag=1;next}/EOF_COMMIT_BODY/{flag=0}flag')
  PR_TITLE=$(echo "$CLAUDE_OUTPUT" | awk '/pr_title<<EOF_PR_TITLE/{flag=1;next}/EOF_PR_TITLE/{flag=0}flag')
  PR_DESCRIPTION=$(echo "$CLAUDE_OUTPUT" | awk '/pr_description<<EOF_PR_DESCRIPTION/{flag=1;next}/EOF_PR_DESCRIPTION/{flag=0}flag')

  export COMMIT_TITLE COMMIT_BODY PR_TITLE PR_DESCRIPTION
}

# Commit changes
commit_changes() {
  if [[ "$HAS_CHANGES" != "true" ]]; then
    if [[ "$EVENT_TYPE" == "openapi-spec-merged" ]]; then
      # Scenario: The managed-service PR was merged, but the spec hasn't changed
      # since our last sync. We still need to add the managed-service commit SHA
      # to the SDK commit message to track which managed-service commit this SDK
      # change corresponds to.
      # Note: This only runs on open SDK PRs - if the SDK PR was already merged,
      # find_sdk_pr would have failed and exited earlier.
      log_info "No changes detected, but adding SHA trailer to existing commit"

      CURRENT_COMMIT_MSG=$(git log -1 --format=%B)

      if echo "$CURRENT_COMMIT_MSG" | grep --quiet "Managed-service-commit-SHA: $MANAGED_SERVICE_SHA"; then
        log_info "SHA trailer already exists in commit"
      else
        NEW_COMMIT_MSG=$(printf "%s\nManaged-service-commit-SHA: %s" "$CURRENT_COMMIT_MSG" "$MANAGED_SERVICE_SHA")
        git commit --amend --message "$NEW_COMMIT_MSG"
      fi
    fi
    return
  fi

  # Stage all generated files and the updated CHANGELOG
  git add internal/spec/openapi.json \
          internal/openapi-generator/api/openapi-modified.json \
          pkg/client/ \
          docs/ \
          README.md \
          CHANGELOG.md

  COMMIT_MSG=$(printf "%s\n\n%s" "$COMMIT_TITLE" "$COMMIT_BODY")

  # Amend the existing commit if updating an SDK PR, create new commit otherwise
  if [[ "$EXISTING_SDK_PR" == "true" ]]; then
    log_info "Amending existing commit"
    git commit --amend --message "$COMMIT_MSG"
  else
    log_info "Creating new commit"
    git commit --message "$COMMIT_MSG"
  fi
}

# Push to fork
# Pushes the SDK branch to the bot's fork. Uses force-with-lease when updating
# existing PRs (since we amend commits), regular push for new branches.
push_to_fork() {
  if [[ "$HAS_CHANGES" != "true" && "$EVENT_TYPE" != "openapi-spec-merged" ]]; then
    return
  fi

  export GIT_ASKPASS="$SCRIPT_DIR/lib/git-askpass.sh"
  export GIT_FORK_USER="$FORK_OWNER"
  export GIT_FORK_PASSWORD="$FORK_PUSH_TOKEN"
  export GIT_TERMINAL_PROMPT=0

  chmod +x "$GIT_ASKPASS"

  # Use force-with-lease for existing PRs (we amended the commit)
  # Use regular push for new branches
  if [[ "$EXISTING_SDK_PR" == "true" ]] || [[ "$EVENT_TYPE" == "openapi-spec-merged" ]]; then
    log_info "Force-pushing to existing branch: $SDK_BRANCH"
    git push --force-with-lease fork "$SDK_BRANCH"
  else
    log_info "Pushing new branch: $SDK_BRANCH"
    git push --set-upstream fork "$SDK_BRANCH"
  fi
}

# Find or create pending deploy branch
# New SDK PRs target an automation/pending-deploy-* branch (not main) to batch multiple
# changes together before release. Reuses the most recent automation/pending-deploy branch
# if one exists, otherwise creates a new one from main.
find_pending_deploy_branch() {
  if [[ "$HAS_CHANGES" != "true" || "$EXISTING_SDK_PR" == "true" ]]; then
    return
  fi

  log_info "Looking for active pending deploy branch..."

  export GH_TOKEN="$CREATE_PR_TOKEN"
  # Find all automation/pending-deploy branches, sorted newest first
  PENDING_BRANCHES=$(gh api "repos/$GITHUB_REPOSITORY/branches" \
    --jq '.[] | select(.name | startswith("automation/pending-deploy-")) | .name' | \
    sort --reverse)

  if [[ -n "$PENDING_BRANCHES" ]]; then
    # Reuse the most recent automation/pending-deploy branch
    PENDING_DEPLOY_BRANCH=$(echo "$PENDING_BRANCHES" | head --lines=1)
    log_info "Found existing pending deploy branch: $PENDING_DEPLOY_BRANCH"
  else
    # No automation/pending-deploy branch exists, create a new one from main
    TIMESTAMP=$(date --utc +%Y%m%d-%H%M%S)
    PENDING_DEPLOY_BRANCH="automation/pending-deploy-$TIMESTAMP"
    log_info "Creating new pending deploy branch: $PENDING_DEPLOY_BRANCH"

    MAIN_SHA=$(gh api "repos/$GITHUB_REPOSITORY/git/ref/heads/main" --jq '.object.sha')
    gh api "repos/$GITHUB_REPOSITORY/git/refs" \
      --method POST \
      --field ref="refs/heads/$PENDING_DEPLOY_BRANCH" \
      --field sha="$MAIN_SHA"
  fi

  export PENDING_DEPLOY_BRANCH
}

# Create or update PR
# Creates a new SDK PR (targeting automation/pending-deploy branch) or updates the
# existing PR's title and description with the latest from Claude.
create_or_update_pr() {
  if [[ "$HAS_CHANGES" != "true" ]]; then
    return
  fi

  export GH_TOKEN="$CREATE_PR_TOKEN"

  if [[ "$EXISTING_SDK_PR" == "true" ]]; then
    # Update the existing PR's title and description
    log_info "Updating existing SDK PR #$SDK_PR_NUMBER"

    gh api --method PATCH "repos/$GITHUB_REPOSITORY/pulls/$SDK_PR_NUMBER" \
      --field title="$PR_TITLE" \
      --field body="$PR_DESCRIPTION"

    log_info "SDK PR updated: https://github.com/$GITHUB_REPOSITORY/pull/$SDK_PR_NUMBER"
  else
    # Create a new PR from the fork to the automation/pending-deploy branch
    log_info "Creating new SDK PR from $FORK_OWNER:$SDK_BRANCH to $GITHUB_REPOSITORY:$PENDING_DEPLOY_BRANCH"

    gh pr create \
      --repo "$GITHUB_REPOSITORY" \
      --head "$FORK_OWNER:$SDK_BRANCH" \
      --base "$PENDING_DEPLOY_BRANCH" \
      --title "$PR_TITLE" \
      --body "$PR_DESCRIPTION"

    log_info "SDK PR created successfully"
  fi
}

# Main execution
main() {
  log_info "=== Starting OpenAPI sync workflow ==="
  log_info "Event type: $EVENT_TYPE"
  log_info "PR URL: $MANAGED_SERVICE_PR_URL"
  log_info "SHA: ${MANAGED_SERVICE_SHA:-<not provided>}"

  check_required_commands gh git jq make curl
  validate_env
  resolve_managed_service_pr
  configure_git
  determine_sdk_branch
  sync_and_regenerate
  update_changelog
  commit_changes
  push_to_fork
  find_pending_deploy_branch
  create_or_update_pr

  if [[ "$HAS_CHANGES" == "false" && "$EVENT_TYPE" != "openapi-spec-merged" ]]; then
    log_info "No changes detected after syncing OpenAPI spec and regenerating client"
  fi

  log_info "=== OpenAPI sync workflow completed ==="
}

main
