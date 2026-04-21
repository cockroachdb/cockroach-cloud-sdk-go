#!/usr/bin/env bash
# Find an SDK PR that corresponds to a managed-service PR URL.
#
# This searches all open PRs in the current repository and inspects
# all commits in each PR to find one containing a trailer:
#   Managed-service-pr-url: <pr_url>
#
# Usage:
#   find_sdk_pr <managed_service_pr_url>
#
# Exports on success:
#   SDK_PR_NUMBER: The SDK PR number
#   SDK_PR_BRANCH: The SDK PR branch name
#
# Returns:
#   0 if found, 1 if not found

# Get the script directory to find helper functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/logging.sh"

find_sdk_pr() {
  local managed_service_pr_url="$1"

  if [[ -z "${managed_service_pr_url:-}" ]]; then
    log_error "managed_service_pr_url is required"
    return 1
  fi

  if [[ -z "${CREATE_PR_TOKEN:-}" ]]; then
    log_error "CREATE_PR_TOKEN environment variable is not set"
    return 1
  fi

  if [[ -z "${GITHUB_REPOSITORY:-}" ]]; then
    log_error "GITHUB_REPOSITORY environment variable is not set"
    return 1
  fi

  log_info "Searching for SDK PR corresponding to: $managed_service_pr_url"

  export GH_TOKEN="$CREATE_PR_TOKEN"

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

        # Get the SDK PR branch
        local sdk_pr_branch
        sdk_pr_branch=$(gh pr view "$sdk_pr_num" --repo "$GITHUB_REPOSITORY" --json headRefName --jq '.headRefName')

        # Export results
        export SDK_PR_NUMBER="$sdk_pr_num"
        export SDK_PR_BRANCH="$sdk_pr_branch"
        return 0
      fi
    done
  done

  log_info "No matching SDK PR found"
  return 1
}
