#!/usr/bin/env bash
set -euo pipefail

# Check deployment status of commits in a pending deploy branch.
#
# This script verifies that all commits in a pending deploy branch
# have been deployed to managed-service by checking their commit trailers
# against the latest release tag.
#
# Required environment variables:
#   PENDING_DEPLOY_BRANCH: The head branch name
#   MANAGED_SERVICE_TOKEN: GitHub token with managed-service read access
#   GITHUB_OUTPUT: Path to GitHub Actions output file
#
# Outputs:
#   Sets has_issues=true/false in GITHUB_OUTPUT
#   Creates result files: not_deployed.txt, missing_trailer.txt, unexpected_status.txt

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
  log_info "=== Starting pending deploy check ==="

  check_required_env "PENDING_DEPLOY_BRANCH" "MANAGED_SERVICE_TOKEN" "GITHUB_OUTPUT" || exit 1

  # Get list of commits in pending deploy branch not yet in main
  get_commits "$PENDING_DEPLOY_BRANCH"

  # Fetch the latest managed-service release tag for comparison
  if ! get_latest_release_tag; then
    log_error "Failed to get latest release tag"
    exit 1
  fi

  # Verify each commit's deployment status and categorize results
  check_all_commits "$LATEST_RELEASE_TAG"

  # Check if any of the result files have content (indicating issues found)
  # These files are used by post-failure-comment.sh to format the PR comment
  if [[ -s not_deployed.txt ]] || [[ -s missing_trailer.txt ]] || [[ -s unexpected_status.txt ]]; then
    log_error "Found commits that are not deployed or potentially not deployed"
    set_output has_issues "true"
  else
    log_info "All commits are deployed"
    set_output has_issues "false"
  fi

  log_info "=== Pending deploy check completed ==="
}

main
