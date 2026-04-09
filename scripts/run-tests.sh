#!/usr/bin/env bash
# Run all bash script tests
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Running all bash script tests..."
echo ""

TOTAL_PASS=0
TOTAL_FAIL=0

for test_file in lib/*-test.sh; do
  if [ -f "$test_file" ]; then
    echo "Running $test_file..."
    if output=$(./"$test_file" 2>&1); then
      echo "$output"
      # Extract pass/fail counts from "Results: X passed, Y failed"
      if [[ "$output" =~ ([0-9]+)\ passed,\ ([0-9]+)\ failed ]]; then
        TOTAL_PASS=$((TOTAL_PASS + ${BASH_REMATCH[1]}))
        TOTAL_FAIL=$((TOTAL_FAIL + ${BASH_REMATCH[2]}))
      fi
    else
      echo "$output"
      TOTAL_FAIL=$((TOTAL_FAIL + 1))
    fi
    echo ""
  fi
done

echo "========================================"
echo "TOTAL: $TOTAL_PASS passed, $TOTAL_FAIL failed"
echo "========================================"

[ "$TOTAL_FAIL" -eq 0 ] || exit 1
