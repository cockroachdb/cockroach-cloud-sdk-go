#!/usr/bin/env bash
# Shared utility functions for pending deploy workflows

# Get the latest managed-service release tag
# Exports: LATEST_RELEASE_TAG
# Returns: 0 on success, 1 on failure
get_latest_release_tag() {
  check_required_env "MANAGED_SERVICE_TOKEN" || return 1

  export GH_TOKEN="$MANAGED_SERVICE_TOKEN"

  log_info "Fetching release tags from managed-service repository"

  # Fetch all release tags matching release-YYYY-MM-DD-N pattern
  # The GitHub API returns results in pages (30 items per page). Without --paginate, we'd only get
  # the first page. With --paginate, gh automatically fetches all pages for us. We need all tags
  # because the API doesn't support sorting by date, so we must fetch everything and sort ourselves.
  local all_tags
  all_tags=$(gh api repos/cockroachlabs/managed-service/tags --paginate --jq '.[].name' | grep --extended-regexp '^release-[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]+$')

  if [[ -z "${all_tags:-}" ]]; then
    log_error "No release tags found in managed-service repository"
    return 1
  fi

  # Sort and get the latest (tags sort lexicographically in chronological order)
  local sorted_tags
  sorted_tags=$(echo "$all_tags" | sort --reverse)
  LATEST_RELEASE_TAG=$(echo "$sorted_tags" | head --lines=1)

  log_info "Latest release tag: $LATEST_RELEASE_TAG"
  export LATEST_RELEASE_TAG
  return 0
}

# Check if a managed-service commit SHA is deployed as of the latest release tag
#
# Args: $1 - managed-service commit SHA
#       $2 - latest release tag
# Returns: 0 if deployed, 1 if not deployed, 2 if uncertain/unexpected
# Outputs to stdout: Status string when deployment status is uncertain
check_deployment_status() {
  local ms_sha="$1"
  local latest_tag="$2"

  if [[ -z "${ms_sha:-}" || -z "${latest_tag:-}" ]]; then
    log_error "Both ms_sha and latest_tag are required"
    echo "missing_parameters"
    return 2
  fi

  if ! check_required_env "MANAGED_SERVICE_TOKEN"; then
    echo "missing_token"
    return 2
  fi

  export GH_TOKEN="$MANAGED_SERVICE_TOKEN"

  # Compare the latest release tag with the commit SHA
  # Status can be: identical, ahead, behind, or diverged
  local compare_status
  if ! compare_status=$(gh api "repos/cockroachlabs/managed-service/compare/${latest_tag}...${ms_sha}" --jq '.status' 2>&1); then
    log_error "Failed to compare commit with release tag: $compare_status"
    echo "unknown"
    return 2
  fi

  case "$compare_status" in
    identical|behind)
      # SHA has been deployed
      return 0
      ;;
    ahead)
      # SHA is ahead of the latest release - not deployed yet
      return 1
      ;;
    *)
      # Unexpected status (e.g., diverged or other)
      echo "$compare_status"
      log_error "Unexpected comparison status: $compare_status"
      return 2
      ;;
  esac
}

# Get commits that are in the pending deploy branch but not in main
get_commits() {
  local branch="$1"

  log_info "Getting commits from $branch not in main"

  # Get commits that are in the pending deploy branch but not in main
  # Format: SHA|subject
  git log origin/main..origin/"$branch" --format="%H|%s" > commits.txt

  if [[ ! -s commits.txt ]]; then
    log_info "No commits found in $branch that are not in main"
    set_output has_issues "false"
    exit 0
  fi

  log_info "Commits to check:"
  cat commits.txt
}

# Check deployment status of all commits
# Creates files categorizing commits by deployment status:
#   - not_deployed.txt: Commits not yet deployed
#   - missing_trailer.txt: Commits missing managed-service SHA trailer
#   - unexpected_status.txt: Commits with unexpected deployment status
check_all_commits() {
  local latest_tag="$1"

  touch not_deployed.txt
  touch missing_trailer.txt
  touch unexpected_status.txt

  # Process each commit
  while IFS='|' read -r sha subject; do
    log_info "Checking commit $sha: $subject"

    # Extract managed-service commit SHA from git commit message trailer
    # Trailers are key-value pairs at the end of commit messages, format:
    # Managed-service-commit-SHA: <sha>
    # This links SDK commits back to the managed-service commit that triggered them
    local ms_sha
    ms_sha=$(git log --max-count=1 --format='%(trailers:key=Managed-service-commit-SHA,valueonly)' "$sha")

    if [[ -z "$ms_sha" ]]; then
      log_info "  No Managed-service-commit-SHA trailer found"
      echo "$sha|$subject" >> missing_trailer.txt
      continue
    fi

    log_info "  Found managed-service SHA: $ms_sha"

    # Check deployment status
    # Return codes: 0=deployed, 1=not deployed, 2=uncertain/unexpected
    # On uncertain status (return 2), the function outputs status details to stdout
    # Note: We use || status=$? to capture non-zero exit codes without triggering set -e
    local output status=0
    output=$(check_deployment_status "$ms_sha" "$latest_tag") || status=$?

    if [[ $status -eq 0 ]]; then
      log_info "  Deployed"
    elif [[ $status -eq 1 ]]; then
      log_info "  Not deployed yet"
      echo "$sha|$subject|$ms_sha" >> not_deployed.txt
    else
      log_error "  Unexpected status: $output"
      echo "$sha|$subject|$ms_sha|$output" >> unexpected_status.txt
    fi
  done < commits.txt
}

# Find the latest pending deploy branch
find_latest_branch() {
  log_info "Looking for pending deploy branches"

  # Find all remote branches matching automation/pending-deploy-YYYYMMDD-HHMMSS
  # The timestamp format ensures lexicographic sorting matches chronological ordering
  # Example: automation/pending-deploy-20250506-143022 (May 6, 2025 at 14:30:22)
  local all_branches
  all_branches=$(git branch --remotes)

  local branches
  branches=$(echo "$all_branches" | grep --extended-regexp 'origin/automation/pending-deploy-[0-9]{8}-[0-9]{6}$' || true)

  if [[ -z "$branches" ]]; then
    log_info "No pending deploy branches found on origin"
    set_output branch ""
    return 0
  fi

  # Find the latest branch (sort --reverse gives newest first due to timestamp format)
  local latest
  latest=$(echo "$branches" | sed 's|^[[:space:]]*origin/||' | sort --reverse | head --lines=1)

  if [[ -z "$latest" ]]; then
    log_error "Failed to parse pending deploy branches"
    exit 1
  fi

  log_info "Latest pending deploy branch: $latest"
  set_output branch "$latest"
  export BRANCH="$latest"
}

# Check if the branch has commits not in main
check_branch_has_new_commits() {
  local branch="$1"

  if [[ -z "$branch" ]]; then
    return 0
  fi

  log_info "Checking for commits in $branch not in main"

  # Count commits in branch but not in origin/main
  local commit_count
  commit_count=$(git rev-list --count origin/main..origin/"$branch")

  if [[ "$commit_count" -eq 0 ]]; then
    log_info "No new commits in $branch"
    set_output has_new_commits "false"
    return 0
  fi

  log_info "Found $commit_count commit(s) in $branch not in main"
  set_output has_new_commits "true"
  export HAS_NEW_COMMITS="true"
}

# Create or check for existing PR
create_pr_if_not_exists() {
  local branch="$1"
  local has_new_commits="$2"

  if [[ -z "$branch" ]] || [[ "$has_new_commits" != "true" ]]; then
    return 0
  fi

  log_info "Checking if PR already exists for $branch"

  # Check if PR already exists
  local existing_pr_url
  existing_pr_url=$(gh pr list --head "$branch" --base main --json url --jq '.[0].url' || echo "")

  if [[ -n "$existing_pr_url" ]] && [[ "$existing_pr_url" != "null" ]]; then
    log_info "PR already exists: $existing_pr_url"
    set_output pr_url "$existing_pr_url"
    return 0
  fi

  log_info "Creating new PR for $branch"

  # Create new PR
  local pr_body
  pr_body="This PR represents commits from \`$branch\` that are pending deployment. These commits were generated in response to Cockroach Cloud API changes in the managed-service repository."

  local pr_url
  pr_url=$(gh pr create \
    --head "$branch" \
    --base main \
    --title "Applying pending deploy changes: $branch" \
    --body "$pr_body")

  if [[ -z "$pr_url" ]]; then
    log_error "Failed to create PR"
    exit 1
  fi

  log_info "Created PR: $pr_url"
  set_output pr_url "$pr_url"
}
