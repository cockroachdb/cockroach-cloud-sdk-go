#!/usr/bin/env bash
# Tests for openapi-sync-helpers.sh functions
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."
source ./test-helpers.sh
source ./lib/actions-helpers.sh
source ./lib/validation.sh
source ./lib/openapi-sync-helpers.sh

# --- parse_managed_service_pr_url tests ---

test_parse_pr_url_valid() {
  MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/1234"
  parse_managed_service_pr_url
  [ "$MS_OWNER" = "cockroachlabs" ] && \
  [ "$MS_REPO" = "managed-service" ] && \
  [ "$MS_PR_NUMBER" = "1234" ]
}
expect_success "parse_managed_service_pr_url: parses valid URL" test_parse_pr_url_valid

test_parse_pr_url_different_org() {
  MANAGED_SERVICE_PR_URL="https://github.com/myorg/myrepo/pull/5678"
  parse_managed_service_pr_url
  [ "$MS_OWNER" = "myorg" ] && \
  [ "$MS_REPO" = "myrepo" ] && \
  [ "$MS_PR_NUMBER" = "5678" ]
}
expect_success "parse_managed_service_pr_url: handles different org/repo" test_parse_pr_url_different_org

test_parse_pr_url_invalid_no_pr() {
  MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects URL without /pull/" test_parse_pr_url_invalid_no_pr

test_parse_pr_url_invalid_domain() {
  MANAGED_SERVICE_PR_URL="https://gitlab.com/cockroachlabs/repo/pull/123"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects non-GitHub domain" test_parse_pr_url_invalid_domain

test_parse_pr_url_invalid_format() {
  MANAGED_SERVICE_PR_URL="not-a-url"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects malformed URL" test_parse_pr_url_invalid_format

test_parse_pr_url_error_message() {
  MANAGED_SERVICE_PR_URL="invalid-url"
  local output
  output=$(parse_managed_service_pr_url 2>&1) || true
  echo "$output" | check_contains "Invalid PR URL format"
}
expect_success "parse_managed_service_pr_url: shows error message" test_parse_pr_url_error_message

test_parse_pr_url_missing_pr_number() {
  MANAGED_SERVICE_PR_URL="https://github.com/owner/repo/pull/"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects missing PR number" test_parse_pr_url_missing_pr_number

# --- validate_env tests ---

test_validate_env_all_present() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    export MANAGED_SERVICE_PR_URL="https://github.com/org/repo/pull/123"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env
  )
}
expect_success "validate_env: succeeds with all vars present" test_validate_env_all_present

test_validate_env_missing_event_type() {
  (
    export MANAGED_SERVICE_PR_URL="https://github.com/org/repo/pull/123"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env 2>&1
  )
}
expect_failure "validate_env: fails when EVENT_TYPE missing" test_validate_env_missing_event_type

test_validate_env_missing_pr_url() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env 2>&1
  )
}
expect_failure "validate_env: fails when MANAGED_SERVICE_PR_URL missing" test_validate_env_missing_pr_url

test_validate_env_merged_event_requires_sha() {
  (
    export EVENT_TYPE="openapi-spec-merged"
    export MANAGED_SERVICE_PR_URL="https://github.com/org/repo/pull/123"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env
  )
}
expect_failure "validate_env: merged event requires MANAGED_SERVICE_SHA" test_validate_env_merged_event_requires_sha

test_validate_env_merged_event_with_sha() {
  (
    export EVENT_TYPE="openapi-spec-merged"
    export MANAGED_SERVICE_PR_URL="https://github.com/org/repo/pull/123"
    export MANAGED_SERVICE_SHA="abc123def456"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env
  )
}
expect_success "validate_env: merged event succeeds with SHA" test_validate_env_merged_event_with_sha

test_validate_env_changed_event_without_sha() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    export MANAGED_SERVICE_PR_URL="https://github.com/org/repo/pull/123"
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    validate_env
  )
}
expect_success "validate_env: changed event doesn't require SHA" test_validate_env_changed_event_without_sha

test_validate_env_error_message() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    # Missing MANAGED_SERVICE_PR_URL
    export MANAGED_SERVICE_TOKEN="token123"
    export FORK_PUSH_TOKEN="token456"
    export FORK_OWNER="bot-user"
    export GH_TOKEN="token789"
    export GITHUB_REPOSITORY="owner/repo"
    local output
    output=$(validate_env 2>&1) || true
    echo "$output" | check_contains "MANAGED_SERVICE_PR_URL"
    echo "$output" | check_contains "required"
  )
}
expect_success "validate_env: shows which var is missing" test_validate_env_error_message

print_results
