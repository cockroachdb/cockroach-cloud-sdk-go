#!/usr/bin/env bash
set -euo pipefail

# Sync OpenAPI spec from managed-service PR and regenerate client.
#
# This script orchestrates the entire OpenAPI sync workflow:
# 1. Validates event payload
# 2. Parses managed-service PR metadata
# 3. Determines head and base branches (finds existing SDK PR or creates new branches)
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
#   GH_TOKEN: GitHub token for gh CLI operations
#   GITHUB_REPOSITORY: Repository in owner/repo format

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"
source "$SCRIPT_DIR/lib/openapi-sync-helpers.sh"

# Check for required commands
check_required_commands "gh" "git" "jq" "make" "curl" || exit 1

# Main execution
main() {
  log_info "=== Starting OpenAPI sync workflow ==="
  log_info "Event type: $EVENT_TYPE"
  log_info "PR URL: $MANAGED_SERVICE_PR_URL"
  log_info "SHA: ${MANAGED_SERVICE_SHA:-<not provided>}"

  validate_env
  parse_managed_service_pr_url
  configure_git
  determine_head_and_base_branches
  sync_and_regenerate
  update_changelog
  commit_changes
  push_to_fork
  create_or_update_pr

  if [[ "$HAS_CHANGES" == "false" && "$EVENT_TYPE" != "openapi-spec-merged" ]]; then
    log_info "No changes detected after syncing OpenAPI spec and regenerating client"
  fi

  log_info "=== OpenAPI sync workflow completed ==="
}

main
