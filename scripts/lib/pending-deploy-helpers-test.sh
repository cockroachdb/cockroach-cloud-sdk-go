#!/usr/bin/env bash
# Tests for pending-deploy-helpers.sh functions
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."
source ./test-helpers.sh
source ./lib/actions-helpers.sh
source ./lib/validation.sh
source ./lib/pending-deploy-helpers.sh

# Create a temporary directory for test artifacts
TEST_DIR=$(mktemp -d)
cd "$TEST_DIR"

cleanup() {
  cd - > /dev/null
  rm -rf "$TEST_DIR"
}
trap cleanup EXIT

# Mock set_output to avoid writing to real GITHUB_OUTPUT
set_output() {
  echo "set_output: $1=$2"
}

# --- check_deployment_status tests ---

test_check_deployment_status_missing_sha() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    local output
    output=$(check_deployment_status "" "release-2024-01-01-1" 2>&1) || true
    echo "$output" | check_contains "missing_parameters"
  )
}
expect_success "check_deployment_status: missing sha parameter outputs missing_parameters" test_check_deployment_status_missing_sha

test_check_deployment_status_missing_tag() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    local output
    output=$(check_deployment_status "abc123" "" 2>&1) || true
    echo "$output" | check_contains "missing_parameters"
  )
}
expect_success "check_deployment_status: missing tag parameter outputs missing_parameters" test_check_deployment_status_missing_tag

test_check_deployment_status_missing_token() {
  (
    unset MANAGED_SERVICE_TOKEN
    local output
    output=$(check_deployment_status "abc123" "release-2024-01-01-1" 2>&1) || true
    echo "$output" | check_contains "missing_token"
  )
}
expect_success "check_deployment_status: missing token outputs missing_token" test_check_deployment_status_missing_token

test_check_deployment_status_identical() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    # Mock gh command to return "identical" status (after jq extraction)
    gh() {
      if [[ "$*" == *"compare"* ]] && [[ "$*" == *"--jq"* ]]; then
        echo "identical"
      fi
    }
    export -f gh
    check_deployment_status "abc123" "release-2024-01-01-1"
  )
}
expect_success "check_deployment_status: identical status returns 0" test_check_deployment_status_identical

test_check_deployment_status_behind() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    # Mock gh command to return "behind" status (after jq extraction)
    gh() {
      if [[ "$*" == *"compare"* ]] && [[ "$*" == *"--jq"* ]]; then
        echo "behind"
      fi
    }
    export -f gh
    check_deployment_status "abc123" "release-2024-01-01-1"
  )
}
expect_success "check_deployment_status: behind status returns 0" test_check_deployment_status_behind

test_check_deployment_status_ahead() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    # Mock gh command to return "ahead" status (after jq extraction)
    gh() {
      if [[ "$*" == *"compare"* ]] && [[ "$*" == *"--jq"* ]]; then
        echo "ahead"
      fi
    }
    export -f gh
    check_deployment_status "abc123" "release-2024-01-01-1" 2>&1
  )
}
expect_failure "check_deployment_status: ahead status returns 1" test_check_deployment_status_ahead

test_check_deployment_status_diverged() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    # Mock gh command to return "diverged" status (after jq extraction)
    gh() {
      if [[ "$*" == *"compare"* ]] && [[ "$*" == *"--jq"* ]]; then
        echo "diverged"
      fi
    }
    export -f gh
    local output exit_code
    output=$(check_deployment_status "abc123" "release-2024-01-01-1" 2>&1) || exit_code=$?
    [[ "$exit_code" -eq 2 ]] && echo "$output" | check_contains "diverged"
  )
}
expect_success "check_deployment_status: diverged status returns 2 and outputs status" test_check_deployment_status_diverged

test_check_deployment_status_api_failure() {
  (
    export MANAGED_SERVICE_TOKEN="token"
    # Mock gh command to fail
    gh() {
      if [[ "$*" == *"compare"* ]]; then
        echo "API error" >&2
        return 1
      fi
    }
    export -f gh
    local output exit_code
    output=$(check_deployment_status "abc123" "release-2024-01-01-1" 2>&1) || exit_code=$?
    [[ "$exit_code" -eq 2 ]] && echo "$output" | check_contains "unknown"
  )
}
expect_success "check_deployment_status: API failure returns 2 and outputs unknown" test_check_deployment_status_api_failure

# --- find_latest_branch tests ---

test_find_latest_branch_no_branches() {
  (
    # Mock git to return no matching branches
    git() {
      if [[ "$*" == "branch --remotes" ]]; then
        echo "  origin/main"
        echo "  origin/feature-branch"
      else
        command git "$@"
      fi
    }
    export -f git

    find_latest_branch 2>&1 | check_contains "No pending deploy branches"
  )
}
expect_success "find_latest_branch: no matching branches found" test_find_latest_branch_no_branches

test_find_latest_branch_single_branch() {
  (
    # Mock git to return one pending deploy branch
    git() {
      if [[ "$*" == "branch --remotes" ]]; then
        echo "  origin/main"
        echo "  origin/automation/pending-deploy-20240501-120000"
      else
        command git "$@"
      fi
    }
    export -f git

    local output
    output=$(find_latest_branch 2>&1)
    echo "$output" | check_contains "automation/pending-deploy-20240501-120000"
  )
}
expect_success "find_latest_branch: single branch found" test_find_latest_branch_single_branch

test_find_latest_branch_multiple_branches_picks_latest() {
  (
    # Mock git to return multiple pending deploy branches
    git() {
      if [[ "$*" == "branch --remotes" ]]; then
        echo "  origin/main"
        echo "  origin/automation/pending-deploy-20240501-120000"
        echo "  origin/automation/pending-deploy-20240502-140000"
        echo "  origin/automation/pending-deploy-20240501-150000"
      else
        command git "$@"
      fi
    }
    export -f git

    local output
    output=$(find_latest_branch 2>&1)
    # Should pick the latest timestamp: 20240502-140000
    echo "$output" | check_contains "automation/pending-deploy-20240502-140000"
  )
}
expect_success "find_latest_branch: picks latest from multiple branches" test_find_latest_branch_multiple_branches_picks_latest

# --- check_all_commits tests ---

test_check_all_commits_missing_trailer() {
  (
    export MANAGED_SERVICE_TOKEN="token"

    # Create commits.txt with one commit
    echo "abc123|Add new feature" > commits.txt

    # Mock git log to return empty trailer
    git() {
      if [[ "$*" == *"trailers"* ]]; then
        echo ""
      else
        command git "$@"
      fi
    }
    export -f git

    check_all_commits "release-2024-01-01-1"

    # Check that commit was added to missing_trailer.txt
    grep --quiet "abc123|Add new feature" missing_trailer.txt
  )
}
expect_success "check_all_commits: commit without trailer goes to missing_trailer.txt" test_check_all_commits_missing_trailer

test_check_all_commits_deployed() {
  (
    export MANAGED_SERVICE_TOKEN="token"

    # Create commits.txt with one commit
    echo "abc123|Add new feature" > commits.txt

    # Mock git log to return a trailer
    git() {
      if [[ "$*" == *"trailers"* ]]; then
        echo "def456"
      else
        command git "$@"
      fi
    }
    export -f git

    # Mock check_deployment_status to return 0 (deployed)
    check_deployment_status() {
      return 0
    }
    export -f check_deployment_status

    check_all_commits "release-2024-01-01-1"

    # Check that commit was NOT added to any error file
    [[ ! -s not_deployed.txt ]] && [[ ! -s unexpected_status.txt ]]
  )
}
expect_success "check_all_commits: deployed commit not in error files" test_check_all_commits_deployed

test_check_all_commits_not_deployed() {
  (
    export MANAGED_SERVICE_TOKEN="token"

    # Create commits.txt with one commit
    echo "abc123|Add new feature" > commits.txt

    # Mock git log to return a trailer
    git() {
      if [[ "$*" == *"trailers"* ]]; then
        echo "def456"
      else
        command git "$@"
      fi
    }
    export -f git

    # Mock check_deployment_status to return 1 (not deployed)
    check_deployment_status() {
      return 1
    }
    export -f check_deployment_status

    check_all_commits "release-2024-01-01-1"

    # Check that commit was added to not_deployed.txt
    grep --quiet "abc123|Add new feature|def456" not_deployed.txt
  )
}
expect_success "check_all_commits: not deployed commit goes to not_deployed.txt" test_check_all_commits_not_deployed

test_check_all_commits_unexpected_status() {
  (
    export MANAGED_SERVICE_TOKEN="token"

    # Create commits.txt with one commit
    echo "abc123|Add new feature" > commits.txt

    # Mock git log to return a trailer
    git() {
      if [[ "$*" == *"trailers"* ]]; then
        echo "def456"
      else
        command git "$@"
      fi
    }
    export -f git

    # Mock check_deployment_status to return 2 (unexpected) with output
    check_deployment_status() {
      echo "diverged"
      return 2
    }
    export -f check_deployment_status

    check_all_commits "release-2024-01-01-1" 2>&1

    # Check that commit was added to unexpected_status.txt with status
    grep --quiet "abc123|Add new feature|def456|diverged" unexpected_status.txt
  )
}
expect_success "check_all_commits: unexpected status goes to unexpected_status.txt" test_check_all_commits_unexpected_status

test_check_all_commits_mixed_statuses() {
  (
    export MANAGED_SERVICE_TOKEN="token"

    # Create commits.txt with multiple commits
    cat > commits.txt <<'EOF'
commit1|First commit
commit2|Second commit
commit3|Third commit
commit4|Fourth commit
EOF

    # Mock git log to return different trailers
    git() {
      if [[ "$*" == *"trailers"* ]]; then
        local sha="$4"
        case "$sha" in
          commit1) echo "" ;;  # No trailer
          commit2) echo "ms-sha-2" ;;
          commit3) echo "ms-sha-3" ;;
          commit4) echo "ms-sha-4" ;;
        esac
      else
        command git "$@"
      fi
    }
    export -f git

    # Mock check_deployment_status to return different statuses
    check_deployment_status() {
      local ms_sha="$1"
      case "$ms_sha" in
        ms-sha-2) return 0 ;;  # Deployed
        ms-sha-3) return 1 ;;  # Not deployed
        ms-sha-4) echo "diverged"; return 2 ;;  # Unexpected
      esac
    }
    export -f check_deployment_status

    check_all_commits "release-2024-01-01-1" 2>&1

    # Verify categorization
    grep --quiet "commit1|First commit" missing_trailer.txt && \
    grep --quiet "commit3|Third commit|ms-sha-3" not_deployed.txt && \
    grep --quiet "commit4|Fourth commit|ms-sha-4|diverged" unexpected_status.txt && \
    ! grep --quiet "commit2" missing_trailer.txt && \
    ! grep --quiet "commit2" not_deployed.txt && \
    ! grep --quiet "commit2" unexpected_status.txt
  )
}
expect_success "check_all_commits: correctly categorizes mixed commit statuses" test_check_all_commits_mixed_statuses

print_results
