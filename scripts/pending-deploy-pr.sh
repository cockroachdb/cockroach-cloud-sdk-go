#!/usr/bin/env bash
set -euo pipefail

# Create or update a PR for the latest pending deploy branch.
#
# This script finds the latest pending deploy branch and creates a PR
# to merge it into main if it has commits that are not yet in main.
#
# Required environment variables:
#   GITHUB_TOKEN: GitHub token for creating PRs
#   GITHUB_REPOSITORY: Repository in owner/repo format
#   GITHUB_OUTPUT: Path to GitHub Actions output file

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"
source "$SCRIPT_DIR/lib/pending-deploy-helpers.sh"

# Check for required commands
check_required_commands "gh" "git" || exit 1

# Main execution
main() {
  log_info "=== Starting pending deploy PR workflow ==="

  check_required_env "GITHUB_TOKEN" "GITHUB_REPOSITORY" "GITHUB_OUTPUT" || exit 1

  # Set GH_TOKEN for gh CLI commands
  export GH_TOKEN="$GITHUB_TOKEN"

  # Find the most recent automation/pending-deploy-* branch
  find_latest_branch

  if [[ -n "${BRANCH:-}" ]]; then
    # Check if the branch has commits not yet in main
    check_branch_has_new_commits "$BRANCH"
    create_pr_if_not_exists "$BRANCH" "${HAS_NEW_COMMITS:-false}"
  fi

  log_info "=== Pending deploy PR workflow completed ==="
}

main
