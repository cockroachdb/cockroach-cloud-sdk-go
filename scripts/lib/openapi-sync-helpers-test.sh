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

test_parse_pr_url_different_repo() {
  MANAGED_SERVICE_PR_URL="https://github.com/myorg/myrepo/pull/123"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects URL from different repo" test_parse_pr_url_different_repo

test_parse_pr_url_no_pr() {
  MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects URL without /pull/" test_parse_pr_url_no_pr

test_parse_pr_url_malformed() {
  MANAGED_SERVICE_PR_URL="not-a-url"
  parse_managed_service_pr_url 2>&1
}
expect_failure "parse_managed_service_pr_url: rejects malformed URL" test_parse_pr_url_malformed

# --- validate_event_payload tests ---

test_validate_event_payload_valid_changed_event() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    export MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/123"
    validate_event_payload
  )
}
expect_success "validate_event_payload: accepts valid changed event" test_validate_event_payload_valid_changed_event

test_validate_event_payload_invalid_event_type() {
  (
    export EVENT_TYPE="invalid-event"
    export MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/123"
    validate_event_payload 2>&1
  )
}
expect_failure "validate_event_payload: rejects invalid EVENT_TYPE" test_validate_event_payload_invalid_event_type

test_validate_event_payload_merged_event_requires_sha() {
  (
    export EVENT_TYPE="openapi-spec-merged"
    export MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/123"
    validate_event_payload 2>&1
  )
}
expect_failure "validate_event_payload: merged event requires MANAGED_SERVICE_SHA" test_validate_event_payload_merged_event_requires_sha

test_validate_event_payload_merged_event_with_sha() {
  (
    export EVENT_TYPE="openapi-spec-merged"
    export MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/123"
    export MANAGED_SERVICE_SHA="abc123def456"
    validate_event_payload
  )
}
expect_success "validate_event_payload: merged event succeeds with SHA" test_validate_event_payload_merged_event_with_sha

test_validate_event_payload_changed_event_rejects_sha() {
  (
    export EVENT_TYPE="openapi-spec-changed"
    export MANAGED_SERVICE_PR_URL="https://github.com/cockroachlabs/managed-service/pull/123"
    export MANAGED_SERVICE_SHA="abc123def456"
    validate_event_payload 2>&1
  )
}
expect_failure "validate_event_payload: changed event rejects SHA" test_validate_event_payload_changed_event_rejects_sha

print_results
