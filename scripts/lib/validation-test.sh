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

print_results
