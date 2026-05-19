#!/usr/bin/env bash
set -euo pipefail

# Post a PR comment with deployment check failure details.
#
# This script formats and posts a comment to a GitHub PR explaining
# which commits in a pending deploy branch have not been deployed yet.
#
# Required environment variables:
#   PR_NUMBER: Pull request number
#   GITHUB_TOKEN: GitHub token for posting PR comments
#
# Required files (created by pending-deploy-check.sh):
#   not_deployed.txt: Commits that are definitely not deployed
#   missing_trailer.txt: Commits missing the managed-service SHA trailer
#   unexpected_status.txt: Commits with unexpected deployment status

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"

# Cleanup function for temporary files
cleanup() {
  rm -f comment.md
}
trap cleanup EXIT

# Check for required commands
check_required_commands "gh" || exit 1

# Validate required environment variables
check_required_env "PR_NUMBER" "GITHUB_TOKEN" || exit 1

# Set GH_TOKEN for gh CLI commands
export GH_TOKEN="$GITHUB_TOKEN"

log_info "Formatting deployment check failure comment for PR #$PR_NUMBER"

# Build comment body
cat > comment.md <<'EOF'
## Pre-Deploy Check Failed

This PR contains commits that are not yet deployed or potentially not deployed in managed-service.

EOF

# Add not deployed section
if [[ -s not_deployed.txt ]]; then
  cat >> comment.md <<'EOF'
### Not Deployed Commits

These commits reference managed-service SHAs that are definitively not deployed:

EOF

  while IFS='|' read -r sha subject ms_sha; do
    cat >> comment.md <<EOF
- $subject: \`$sha\`
  - Associated managed-service SHA: \`$ms_sha\`

EOF
  done < not_deployed.txt
fi

# Add potentially not deployed section
if [[ -s missing_trailer.txt ]] || [[ -s unexpected_status.txt ]]; then
  cat >> comment.md <<'EOF'
### Potentially Not Deployed Commits

EOF

  # Missing trailers section
  if [[ -s missing_trailer.txt ]]; then
    cat >> comment.md <<'EOF'
#### Missing Trailers

These commits lack a managed-service SHA trailer:

EOF

    while IFS='|' read -r sha subject; do
      cat >> comment.md <<EOF
- $subject: \`$sha\`

EOF
    done < missing_trailer.txt
  fi

  # Unexpected status section
  if [[ -s unexpected_status.txt ]]; then
    cat >> comment.md <<'EOF'
#### Unexpected Deploy Status

These commits have unexpected deployment status:

EOF

    while IFS='|' read -r sha subject ms_sha output; do
      cat >> comment.md <<EOF
- $subject: \`$sha\`
  - Associated managed-service SHA: \`$ms_sha\`
  - Status: \`$output\`

EOF
    done < unexpected_status.txt
  fi
fi

cat >> comment.md <<'EOF'

---

**Action Required:** This PR is blocked from merging until all commits are confirmed deployed in managed-service.

- For commits with missing trailers: Add the `Managed-service-commit-SHA:` trailer to the commit message
- For not deployed commits: Remove commit from pending deploy branch and wait for the corresponding managed-service release
- This check will re-run automatically when the PR is updated
EOF

# Post comment
gh pr comment "$PR_NUMBER" --body-file comment.md

log_info "Posted deployment failure comment to PR #$PR_NUMBER"
