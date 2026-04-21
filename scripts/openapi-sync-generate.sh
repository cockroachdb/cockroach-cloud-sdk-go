#!/usr/bin/env bash
set -euo pipefail

# Generate updated SDK client from OpenAPI spec.
#
# This script:
# 1. Validates event payload
# 2. Parses managed-service PR metadata
# 3. Configures Git
# 4. Determines head and base branches (finds existing SDK PR or creates new branches)
# 5. Fetches OpenAPI spec and regenerates client code
#
# Environment variables:
#   EVENT_TYPE: Type of event triggering the sync (openapi-spec-changed or openapi-spec-merged)
#   MANAGED_SERVICE_PR_URL: URL of the managed-service PR containing spec changes
#   MANAGED_SERVICE_SHA: Commit SHA in managed-service to fetch spec from (merged events only)
#   MANAGED_SERVICE_TOKEN: GitHub token with read access to managed-service repo (secret)
#   FORK_PUSH_TOKEN: GitHub token with push access to bot fork (secret)
#   FORK_OWNER: GitHub username of the bot fork owner
#   GH_TOKEN: GitHub token for gh CLI operations (secret)
#   GITHUB_REPOSITORY: Current repository in owner/repo format

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"
source "$SCRIPT_DIR/lib/openapi-sync-helpers.sh"

# Main execution
main() {
  log_info "=== Starting OpenAPI sync step: Generate ==="
  log_info "Event type: $EVENT_TYPE"
  log_info "PR URL: $MANAGED_SERVICE_PR_URL"
  log_info "SHA: ${MANAGED_SERVICE_SHA:-<not provided>}"

  check_required_commands "gh" "git" "jq" "make" || exit 1
  check_required_env "EVENT_TYPE" "MANAGED_SERVICE_PR_URL" "MANAGED_SERVICE_TOKEN" \
    "FORK_PUSH_TOKEN" "FORK_OWNER" "GH_TOKEN" "GITHUB_REPOSITORY" || exit 1
  validate_event_payload
  parse_managed_service_pr_url
  configure_git
  determine_head_and_base_branches
  sync_and_regenerate

  # Export outputs for subsequent steps
  {
    echo "base_branch=$BASE_BRANCH"
    echo "has_changes=$HAS_CHANGES"
    echo "existing_sdk_pr=$EXISTING_SDK_PR"
    echo "head_branch=$HEAD_BRANCH"
    echo "sdk_pr_number=${SDK_PR_NUMBER:-}"
  } >> "$GITHUB_OUTPUT"

  log_info "=== OpenAPI sync step completed: Generate ==="
}

main
