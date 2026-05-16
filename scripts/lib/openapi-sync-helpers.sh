#!/usr/bin/env bash
# Helper functions for OpenAPI sync workflow

# Set up git authentication using GIT_ASKPASS
setup_git_auth() {
  local user="$1"
  local password="$2"
  export GIT_ASKPASS="$SCRIPT_DIR/git-askpass.sh"
  export GIT_USER="$user"
  export GIT_PASSWORD="$password"
  export GIT_TERMINAL_PROMPT=0
}

# Find an SDK PR that corresponds to a managed-service PR URL.
#
# This searches all open PRs in the current repository and inspects
# all commits in each PR to find one containing a trailer:
#   Managed-service-pr-url: <pr_url>
#
# Args: $1 - managed-service PR URL
# Exports on success:
#   SDK_PR_NUMBER: The SDK PR number
#   HEAD_BRANCH: The SDK PR head branch name
#   BASE_BRANCH: The SDK PR base branch name
# Returns: 0 if found, 1 if not found
find_sdk_pr() {
  local managed_service_pr_url="$1"

  if [[ -z "${managed_service_pr_url:-}" ]]; then
    log_error "managed_service_pr_url is required"
    return 1
  fi

  if [[ -z "${GH_TOKEN:-}" ]]; then
    log_error "GH_TOKEN environment variable is required"
    return 1
  fi

  if [[ -z "${GITHUB_REPOSITORY:-}" ]]; then
    log_error "GITHUB_REPOSITORY environment variable is required"
    return 1
  fi

  log_info "Searching for SDK PR corresponding to: $managed_service_pr_url"

  # Get all open SDK PRs
  local sdk_open_prs
  sdk_open_prs=$(gh pr list --repo "$GITHUB_REPOSITORY" --state open --json number --jq '.[].number')

  for sdk_pr_num in $sdk_open_prs; do
    log_info "Checking SDK PR #$sdk_pr_num"

    # Get all commits in this SDK PR
    local sdk_commits
    sdk_commits=$(gh api "repos/$GITHUB_REPOSITORY/pulls/$sdk_pr_num/commits" --jq '.[].sha')

    # Check each commit's message for the trailer
    for sdk_commit_sha in $sdk_commits; do
      local sdk_commit_msg
      sdk_commit_msg=$(gh api "repos/$GITHUB_REPOSITORY/git/commits/$sdk_commit_sha" --jq '.message')

      if echo "$sdk_commit_msg" | grep --quiet --fixed-strings "Managed-service-pr-url: $managed_service_pr_url"; then
        log_info "Found matching SDK PR #$sdk_pr_num (commit $sdk_commit_sha)"

        # Get the SDK PR details
        local pr_data
        pr_data=$(gh pr view "$sdk_pr_num" --repo "$GITHUB_REPOSITORY" --json headRefName,baseRefName)

        # Export results
        export SDK_PR_NUMBER="$sdk_pr_num"
        export HEAD_BRANCH=$(echo "$pr_data" | jq -r '.headRefName')
        export BASE_BRANCH=$(echo "$pr_data" | jq -r '.baseRefName')
        return 0
      fi
    done
  done

  log_info "No matching SDK PR found"
  return 1
}

# Validate required environment variables
validate_env() {
  check_required_env "EVENT_TYPE" "MANAGED_SERVICE_PR_URL" "MANAGED_SERVICE_TOKEN" "FORK_PUSH_TOKEN" "FORK_OWNER" "GH_TOKEN" "GITHUB_REPOSITORY" || exit 1

  # Validate EVENT_TYPE is one of the allowed values
  if [[ "$EVENT_TYPE" != "openapi-spec-changed" && "$EVENT_TYPE" != "openapi-spec-merged" ]]; then
    log_error "Invalid EVENT_TYPE: $EVENT_TYPE. Must be 'openapi-spec-changed' or 'openapi-spec-merged'"
    exit 1
  fi

  if [[ "$EVENT_TYPE" == "openapi-spec-merged" && -z "${MANAGED_SERVICE_SHA:-}" ]]; then
    log_error "MANAGED_SERVICE_SHA is required for openapi-spec-merged event"
    exit 1
  fi

  if [[ "$EVENT_TYPE" == "openapi-spec-changed" && -n "${MANAGED_SERVICE_SHA:-}" ]]; then
    log_error "Input 'sha' should not be provided for openapi-spec-changed event"
    exit 1
  fi
}

# Parse managed-service PR metadata from URL
#
# Exports:
#   MS_OWNER: The GitHub organization/user that owns the managed-service repo
#   MS_REPO: The managed-service repository name
#   MS_PR_NUMBER: The managed-service PR number
parse_managed_service_pr_url() {
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
  git config user.name "$FORK_OWNER"
  git config user.email "$FORK_OWNER@users.noreply.github.com"

  FORK_URL="https://github.com/$FORK_OWNER/$(basename "$GITHUB_REPOSITORY").git"
  log_info "Adding fork remote: $FORK_URL"
  git remote add fork "$FORK_URL"

  # Get the default branch name
  local default_branch
  default_branch=$(gh repo view "$GITHUB_REPOSITORY" --json defaultBranchRef --jq '.defaultBranchRef.name')
  log_info "Default branch: $default_branch"

  # Sync fork's default branch with upstream to include any changes from upstream
  # that the fork doesn't have. This prevents feature branches from appearing to add
  # files requiring special permissions (e.g., workflows), which would cause permission issues
  log_info "Syncing fork's $default_branch with origin/$default_branch"

  # Set up askpass for fork push
  setup_git_auth "$FORK_OWNER" "$FORK_PUSH_TOKEN"

  git push --force fork "origin/$default_branch:$default_branch"
}

# Determine head and base branches for the SDK PR
# For existing SDK PRs: uses the PR's head branch and base branch
# For new SDK PRs: finds/creates a base branch, then creates a new head branch from it
# on openapi-spec-changed: may create a new head branch if no SDK PR exists yet
# on openapi-spec-merged: the SDK PR must already exist
#
# Exports:
#   HEAD_BRANCH: The SDK PR head branch name
#   BASE_BRANCH: The SDK PR base branch name (target branch)
#   EXISTING_SDK_PR: "true" if updating existing PR, "false" if creating new PR
determine_head_and_base_branches() {
  # Search all open PRs for one with a commit containing the managed-service PR URL
  if find_sdk_pr "$MANAGED_SERVICE_PR_URL"; then
    log_info "Using existing SDK PR branch: $HEAD_BRANCH (targets $BASE_BRANCH)"

    # Set up askpass for fork operations
    setup_git_auth "$FORK_OWNER" "$FORK_PUSH_TOKEN"

    if ! git fetch fork "$HEAD_BRANCH" 2>&1; then
      log_error "Could not fetch branch $HEAD_BRANCH from fork"
      exit 1
    fi
    git checkout -B "$HEAD_BRANCH" "fork/$HEAD_BRANCH"

    # Fetch the base branch from origin so Claude can diff against the latest state
    setup_git_auth "x-access-token" "$GH_TOKEN"
    git fetch origin "$BASE_BRANCH"

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
    # First, find or create the base branch (pending deploy branch)
    find_pending_deploy_branch
    BASE_BRANCH="$PENDING_DEPLOY_BRANCH"

    TIMESTAMP=$(date --utc +%Y%m%d-%H%M%S)
    HEAD_BRANCH="automation/openapi-spec-change-$TIMESTAMP"
    log_info "Creating new head branch: $HEAD_BRANCH from $BASE_BRANCH"

    # Set up askpass for origin operations
    setup_git_auth "x-access-token" "$GH_TOKEN"

    git fetch origin "$BASE_BRANCH"
    git checkout -b "$HEAD_BRANCH" "origin/$BASE_BRANCH"

    EXISTING_SDK_PR="false"
  fi

  export HEAD_BRANCH BASE_BRANCH EXISTING_SDK_PR
}

# Sync OpenAPI spec and regenerate client
# Fetches the spec from managed-service and regenerates SDK client code.
# For openapi-spec-changed events: fetch from PR branch
# For openapi-spec-merged events: fetch from the exact merged commit SHA
#
# Exports:
#   HAS_CHANGES: "true" if spec/client changed, "false" otherwise
sync_and_regenerate() {
  SPEC_PATH_IN_MS="console/ui/cc-console/assets/docs/api/latest/openapi.json"
  LOCAL_SPEC_PATH="internal/spec/openapi.json"
  LOCAL_SPEC_PATH_NEW="$LOCAL_SPEC_PATH.new"

  # Use commit SHA if provided (merged event), otherwise use PR head (changed event)
  if [[ -n "${MANAGED_SERVICE_SHA:-}" ]]; then
    REF="$MANAGED_SERVICE_SHA"
    log_info "Fetching openapi.json from $MS_OWNER/$MS_REPO at commit $MANAGED_SERVICE_SHA"
  else
    REF="refs/pull/$MS_PR_NUMBER/head"
    log_info "Fetching openapi.json from $MS_OWNER/$MS_REPO PR #$MS_PR_NUMBER"
  fi

  # Fetch the OpenAPI spec file from managed-service
  GH_TOKEN="$MANAGED_SERVICE_TOKEN" gh api "repos/$MS_OWNER/$MS_REPO/contents/$SPEC_PATH_IN_MS?ref=$REF" \
    --jq '.content' | base64 --decode > "$LOCAL_SPEC_PATH_NEW"

  # Validate the downloaded file is valid JSON
  if ! jq empty "$LOCAL_SPEC_PATH_NEW"; then
    log_error "Downloaded spec file is not valid JSON"
    rm "$LOCAL_SPEC_PATH_NEW"
    exit 1
  fi

  # Compare with current spec to see if anything changed
  if cmp --silent "$LOCAL_SPEC_PATH" "$LOCAL_SPEC_PATH_NEW"; then
    log_info "No changes to $LOCAL_SPEC_PATH"
    rm "$LOCAL_SPEC_PATH_NEW"
    HAS_CHANGES="false"
  else
    # Spec changed - replace it and regenerate the client code
    log_info "Spec file changed, replacing $LOCAL_SPEC_PATH"
    mv "$LOCAL_SPEC_PATH_NEW" "$LOCAL_SPEC_PATH"

    # Regenerate SDK client from the new spec using Docker
    # Note: sudo is required because docker containers run as root by default,
    # creating files owned by root. We then chown them back to the current user.
    log_info "Running make generate-openapi-client"
    sudo make generate-openapi-client
    sudo chown --recursive "$(id --user):$(id --group)" internal/openapi-generator/ pkg/client/ docs/ README.md

    # Validate the generated code compiles
    log_info "Validating generated code..."
    if ! make validate; then
      log_error "Generated code failed validation"
      exit 1
    fi

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
#
# Exports:
#   COMMIT_TITLE: Generated commit title
#   COMMIT_BODY: Generated commit body
#   PR_TITLE: Generated PR title
#   PR_DESCRIPTION: Generated PR description
update_changelog() {
  if [[ "$HAS_CHANGES" != "true" ]]; then
    return
  fi

  # Reset CHANGELOG.md to base branch to avoid stale entries in the diff
  # This ensures Claude only sees NEW changes when analyzing git diff
  # Claude will then generate fresh changelog entries based on the actual changes
  log_info "Resetting CHANGELOG.md to base branch to remove stale entries"
  git checkout "origin/$BASE_BRANCH" -- CHANGELOG.md

  log_info "Setting up Claude CLI..."
  curl --fail --silent --show-error --location https://claude.ai/install.sh | bash -s -- "latest"
  log_info "Claude CLI installed: $(claude --version)"

  # Check for required commands
  check_required_commands "claude" || return 1

  # Configure Claude to use Vertex AI for authentication
  log_info "Invoking Claude to update changelog..."
  export CLAUDE_CODE_USE_VERTEX="1"
  export ANTHROPIC_VERTEX_PROJECT_ID="vertex-model-runners"
  export CLOUD_ML_REGION="us-east5"
  export BASE_BRANCH

  local prompt_template="$SCRIPT_DIR/../.claude/skills/generate-sync-metadata.md"

  if [[ ! -f "$prompt_template" ]]; then
    log_error "Prompt template not found at $prompt_template"
    return 1
  fi

  # Create temp files
  local prompt_file=$(mktemp)
  local output_file=$(mktemp)
  trap "rm -f \"$prompt_file\" \"$output_file\"" EXIT RETURN

  # Copy prompt template to temp file
  cp "$prompt_template" "$prompt_file"

  # Call Claude CLI
  log_info "Calling Claude CLI..."
  claude --print \
    --model claude-opus-4-6 \
    --output-format json \
    --max-turns 15 \
    --allowedTools "Read,Edit,Bash(git:*),Bash(printenv:*)" \
    < "$prompt_file" \
    > "$output_file" 2>&1 || {
      log_error "Claude CLI failed"
      cat "$output_file" >&2
      return 1
    }

  # Extract result text from JSON
  local result_text
  result_text=$(jq --raw-output '.result // empty' "$output_file")

  if [[ -z "${result_text:-}" ]]; then
    log_error "Claude produced empty result"
    cat "$output_file" >&2
    return 1
  fi

  log_info "Claude response received, extracting metadata..."

  # Extract commit title (between ---COMMIT_TITLE--- and ---COMMIT_BODY---)
  COMMIT_TITLE=$(echo "$result_text" | awk '/---COMMIT_TITLE---/{flag=1;next}/---COMMIT_BODY---/{flag=0}flag' | grep --invert-match '^[[:space:]]*$' | head --lines=1)

  # Extract commit body (between ---COMMIT_BODY--- and ---PR_TITLE---)
  COMMIT_BODY=$(echo "$result_text" | awk '/---COMMIT_BODY---/{flag=1;next}/---PR_TITLE---/{flag=0}flag')

  # Extract PR title (between ---PR_TITLE--- and ---PR_DESCRIPTION---)
  PR_TITLE=$(echo "$result_text" | awk '/---PR_TITLE---/{flag=1;next}/---PR_DESCRIPTION---/{flag=0}flag' | grep --invert-match '^[[:space:]]*$' | head --lines=1)

  # Extract PR description (after ---PR_DESCRIPTION---)
  PR_DESCRIPTION=$(echo "$result_text" | awk '/---PR_DESCRIPTION---/{flag=1;next}flag')

  # Validate extractions
  if [[ -z "${COMMIT_TITLE:-}" ]]; then
    log_error "Could not extract commit title from Claude output"
    echo "Claude output:" >&2
    echo "$result_text" >&2
    return 1
  fi

  if [[ -z "${PR_TITLE:-}" ]]; then
    log_error "Could not extract PR title from Claude output"
    echo "Claude output:" >&2
    echo "$result_text" >&2
    return 1
  fi

  log_info "Metadata extracted successfully"

  # Export results
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
      git commit --amend --no-edit \
        --trailer "Managed-service-commit-SHA: ${MANAGED_SERVICE_SHA}"
    fi
    return
  fi

  # Stage all generated files and the updated CHANGELOG
  git add internal/spec/openapi.json \
          internal/openapi-generator/api/openapi-modified.json \
          internal/openapi-generator/api/openapi.yaml \
          internal/openapi-generator/.openapi-generator/FILES \
          pkg/client/ \
          docs/ \
          README.md \
          CHANGELOG.md

  # Build base commit message (title + body)
  COMMIT_MSG=$(printf "%s\n\n%s" "$COMMIT_TITLE" "$COMMIT_BODY")

  # For existing PRs: amend the previous OpenAPI sync commit
  # For new PRs: create a new commit on top of the base branch
  if [[ "$EXISTING_SDK_PR" == "true" ]]; then
    log_info "Amending existing commit"
    git commit --amend --message "$COMMIT_MSG" \
      --trailer "Managed-service-pr-url: ${MANAGED_SERVICE_PR_URL}"
  else
    log_info "Creating new commit"
    git commit --message "$COMMIT_MSG" \
      --trailer "Managed-service-pr-url: ${MANAGED_SERVICE_PR_URL}"
  fi

  # Add SHA trailer if provided
  if [[ -n "${MANAGED_SERVICE_SHA:-}" ]]; then
    git commit --amend --no-edit \
      --trailer "Managed-service-commit-SHA: ${MANAGED_SERVICE_SHA}"
  fi
}

# Push to fork
# Pushes the head branch to the bot's fork. Uses force-with-lease when updating
# existing PRs (since we amend commits), regular push for new branches.
push_to_fork() {
  if [[ "$HAS_CHANGES" != "true" && "$EVENT_TYPE" != "openapi-spec-merged" ]]; then
    return
  fi

  setup_git_auth "$FORK_OWNER" "$FORK_PUSH_TOKEN"

  # Use force-with-lease for existing PRs (we amended the commit)
  # Use regular push for new branches
  if [[ "$EXISTING_SDK_PR" == "true" ]] || [[ "$EVENT_TYPE" == "openapi-spec-merged" ]]; then
    log_info "Force-pushing to existing branch: $HEAD_BRANCH"
    git push --force-with-lease fork "$HEAD_BRANCH"
  else
    log_info "Pushing new branch: $HEAD_BRANCH"
    git push --set-upstream fork "$HEAD_BRANCH"
  fi
}

# Find or create pending deploy branch
# New SDK PRs target an automation/pending-deploy-* branch (not main) to batch multiple
# changes together before release. Reuses the most recent automation/pending-deploy branch
# if one exists, otherwise creates a new one from main.
#
# Exports:
#   PENDING_DEPLOY_BRANCH: The pending deploy branch name
find_pending_deploy_branch() {
  log_info "Looking for active pending deploy branch..."

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
# Creates a new SDK PR targeting the base branch or updates an existing PR's
# title and description with the latest from Claude.
create_or_update_pr() {
  if [[ "$HAS_CHANGES" != "true" ]]; then
    return
  fi

  if [[ "$EXISTING_SDK_PR" == "true" ]]; then
    # Update the existing PR's title and description
    log_info "Updating existing SDK PR #$SDK_PR_NUMBER"

    gh api --method PATCH "repos/$GITHUB_REPOSITORY/pulls/$SDK_PR_NUMBER" \
      --field title="$PR_TITLE" \
      --field body="$PR_DESCRIPTION"

    log_info "SDK PR updated: https://github.com/$GITHUB_REPOSITORY/pull/$SDK_PR_NUMBER"
  else
    # Create a new PR from the fork to the base branch
    log_info "Creating new SDK PR from $FORK_OWNER:$HEAD_BRANCH to $GITHUB_REPOSITORY:$BASE_BRANCH"

    gh pr create \
      --repo "$GITHUB_REPOSITORY" \
      --head "$FORK_OWNER:$HEAD_BRANCH" \
      --base "$BASE_BRANCH" \
      --title "$PR_TITLE" \
      --body "$PR_DESCRIPTION"

    log_info "SDK PR created successfully"
  fi
}
