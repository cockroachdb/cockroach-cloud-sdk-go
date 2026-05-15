#!/usr/bin/env bash
# Tests for validation.sh
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."
source ./test-helpers.sh
source ./lib/actions-helpers.sh
source ./lib/validation.sh

# --- check_required_commands tests ---

test_check_required_commands_all_found() {
  check_required_commands bash echo grep
}
expect_success "check_required_commands: all commands found" test_check_required_commands_all_found

test_check_required_commands_one_missing() {
  check_required_commands bash nonexistent_cmd_xyz grep
}
expect_failure "check_required_commands: fails when command missing" test_check_required_commands_one_missing

test_check_required_commands_multiple_missing() {
  local output
  output=$(check_required_commands bash fake_cmd_1 fake_cmd_2 grep 2>&1) || true
  # Should report missing commands
  echo "$output" | check_contains "fake_cmd_1"
  echo "$output" | check_contains "fake_cmd_2"
  # Should NOT report found commands
  ! echo "$output" | check_contains "grep"
}
expect_success "check_required_commands: reports multiple missing" test_check_required_commands_multiple_missing

test_check_required_commands_error_message() {
  local output
  output=$(check_required_commands nonexistent_xyz 2>&1) || true
  echo "$output" | check_contains "::error::"
  echo "$output" | check_contains "not found in PATH"
  echo "$output" | check_contains "nonexistent_xyz"
}
expect_success "check_required_commands: error message format" test_check_required_commands_error_message

test_check_required_commands_empty() {
  check_required_commands
}
expect_success "check_required_commands: handles no arguments" test_check_required_commands_empty

# --- check_required_env tests ---

test_check_required_env_all_set() {
  (
    export TEST_VAR_1="value1"
    export TEST_VAR_2="value2"
    export TEST_VAR_3="value3"
    check_required_env TEST_VAR_1 TEST_VAR_2 TEST_VAR_3
  )
}
expect_success "check_required_env: all env vars set" test_check_required_env_all_set

test_check_required_env_one_missing() {
  (
    export TEST_VAR_1="value1"
    unset TEST_VAR_2
    export TEST_VAR_3="value3"
    check_required_env TEST_VAR_1 TEST_VAR_2 TEST_VAR_3
  )
}
expect_failure "check_required_env: fails when env var missing" test_check_required_env_one_missing

test_check_required_env_multiple_missing() {
  (
    export TEST_VAR_1="value1"
    unset TEST_VAR_2
    unset TEST_VAR_3
    local output
    output=$(check_required_env TEST_VAR_1 TEST_VAR_2 TEST_VAR_3 2>&1) || true
    # Should report missing env vars
    echo "$output" | check_contains "TEST_VAR_2"
    echo "$output" | check_contains "TEST_VAR_3"
    # Should NOT report set env vars
    ! echo "$output" | check_contains "TEST_VAR_1"
  )
}
expect_success "check_required_env: reports multiple missing" test_check_required_env_multiple_missing

test_check_required_env_error_message() {
  (
    unset NONEXISTENT_VAR_XYZ
    local output
    output=$(check_required_env NONEXISTENT_VAR_XYZ 2>&1) || true
    echo "$output" | check_contains "::error::"
    echo "$output" | check_contains "not set"
    echo "$output" | check_contains "NONEXISTENT_VAR_XYZ"
  )
}
expect_success "check_required_env: error message format" test_check_required_env_error_message

test_check_required_env_empty_value() {
  (
    export TEST_VAR_EMPTY=""
    check_required_env TEST_VAR_EMPTY
  )
}
expect_failure "check_required_env: fails when env var is empty string" test_check_required_env_empty_value

test_check_required_env_empty() {
  check_required_env
}
expect_success "check_required_env: handles no arguments" test_check_required_env_empty

print_results
