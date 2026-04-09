#!/usr/bin/env bash
# Minimal test helpers for bash script tests.

PASS=0
FAIL=0

# Assert that input contains a given substring.
# Usage: check_contains "expected" "$file"   (reads file)
#        cmd | check_contains "expected"     (reads stdin)
check_contains() {
  if [ $# -ge 2 ]; then
    grep --quiet --fixed-strings -- "$1" "$2"
  else
    grep --quiet --fixed-strings -- "$1"
  fi
}

_run_test() {
  local name="$1"
  local expected_exit="$2"
  local expected_output="$3"
  shift 3

  local output exit_code
  output=$("$@" 2>&1) && exit_code=0 || exit_code=$?

  if [ "$expected_exit" = "nonzero" ]; then
    if [ "$exit_code" -eq 0 ]; then
      echo "FAIL: $name — expected non-zero exit, got 0"
      echo "  output: $output"
      FAIL=$((FAIL + 1))
      return
    fi
  elif [ "$exit_code" -ne "$expected_exit" ]; then
    echo "FAIL: $name — expected exit $expected_exit, got $exit_code"
    echo "  output: $output"
    FAIL=$((FAIL + 1))
    return
  fi

  if [ -n "$expected_output" ] && ! printf '%s\n' "$output" | check_contains "$expected_output"; then
    echo "FAIL: $name — expected output containing: $expected_output"
    echo "  actual: $output"
    FAIL=$((FAIL + 1))
    return
  fi

  echo "PASS: $name"
  PASS=$((PASS + 1))
}

# expect_success "test name" command [args...]
# Asserts the command exits 0.
expect_success() {
  local name="$1"; shift
  _run_test "$name" 0 "" "$@"
}

# expect_failure "test name" command [args...]
# Asserts the command exits non-zero.
expect_failure() {
  local name="$1"; shift
  _run_test "$name" "nonzero" "" "$@"
}

print_results() {
  echo ""
  echo "Results: $PASS passed, $FAIL failed"
  [ "$FAIL" -eq 0 ] || exit 1
}
