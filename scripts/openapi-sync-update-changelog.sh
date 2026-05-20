#!/usr/bin/env bash
set -euo pipefail

# Update changelog with Claude.
#
# This script uses Claude AI to analyze the git diff, update CHANGELOG.md,
# and generate commit/PR metadata (title, body, description).
#
# Environment variables:
#   BASE_BRANCH: Base branch to diff against
#   HAS_CHANGES: Whether there are changes to process

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"
source "$SCRIPT_DIR/lib/openapi-sync-helpers.sh"

# Main execution
main() {
  log_info "=== Starting OpenAPI sync step: Update changelog ==="

  check_required_commands "curl" || exit 1
  check_required_env "BASE_BRANCH" "HAS_CHANGES" || exit 1
  update_changelog

  # Export outputs for subsequent steps
  echo "commit_title=${COMMIT_TITLE:-}" >> "$GITHUB_OUTPUT"

  {
    echo "commit_body<<EOF"
    echo "${COMMIT_BODY:-}"
    echo "EOF"
  } >> "$GITHUB_OUTPUT"

  echo "pr_title=${PR_TITLE:-}" >> "$GITHUB_OUTPUT"

  {
    echo "pr_description<<EOF"
    echo "${PR_DESCRIPTION:-}"
    echo "EOF"
  } >> "$GITHUB_OUTPUT"

  log_info "=== OpenAPI sync step completed: Update changelog ==="
}

main
