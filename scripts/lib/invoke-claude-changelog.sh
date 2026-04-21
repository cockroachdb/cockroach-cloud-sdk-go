#!/bin/bash
set -euo pipefail

# Invoke Claude to update changelog and generate commit/PR metadata.
#
# This script calls the Claude CLI to analyze code changes, update CHANGELOG.md,
# and return structured metadata for commit messages and PR descriptions.
#
# Usage:
#   invoke-claude-changelog.sh <managed_service_pr_url> [managed_service_sha]
#
# Outputs:
#   commit_title, commit_body, pr_title, pr_description to stdout
#   Exit code 0 on success, 1 on failure

# Get the script directory to find helper functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/logging.sh"

MANAGED_SERVICE_PR_URL="$1"
MANAGED_SERVICE_SHA="${2:-}"

if [[ -z "${MANAGED_SERVICE_PR_URL:-}" ]]; then
  log_error "managed_service_pr_url is required"
  exit 1
fi

# Check that claude CLI is available
if ! command -v claude >/dev/null 2>&1; then
  log_error "claude CLI is not installed or not in PATH"
  exit 1
fi

log_info "Invoking Claude to analyze changes and update changelog..."

PROMPT_TEMPLATE="$SCRIPT_DIR/../../.claude/skills/update-changelog.md"

if [[ ! -f "$PROMPT_TEMPLATE" ]]; then
  log_error "Prompt template not found at $PROMPT_TEMPLATE"
  exit 1
fi

# Export environment variables for Claude to access
export MANAGED_SERVICE_PR_URL
export MANAGED_SERVICE_SHA

# Copy prompt template to temp file
PROMPT_FILE=$(mktemp)
cp "$PROMPT_TEMPLATE" "$PROMPT_FILE"

# Add SHA trailer instruction if provided
if [[ -n "${MANAGED_SERVICE_SHA:-}" ]]; then
  cat >> "$PROMPT_FILE" <<'EOF'

IMPORTANT: The commit body must also include this trailer immediately after the Managed-service-pr-url trailer:
Managed-service-commit-SHA: $MANAGED_SERVICE_SHA
EOF
fi

# Output file for Claude's JSON response
OUTPUT_FILE=$(mktemp)

# Call Claude CLI
log_info "Calling Claude CLI..."
claude --print \
  --model claude-opus-4-6 \
  --output-format json \
  --max-turns 5 \
  --allowedTools "Read,Grep,Glob,Edit,Bash(git:*),Bash(printenv:*)" \
  < "$PROMPT_FILE" \
  > "$OUTPUT_FILE" 2>&1 || {
    log_error "Claude CLI failed"
    cat "$OUTPUT_FILE" >&2
    rm --force "$PROMPT_FILE" "$OUTPUT_FILE"
    exit 1
  }

# Extract result text from JSON
RESULT_TEXT=$(jq --raw-output '.result // empty' "$OUTPUT_FILE")

if [[ -z "${RESULT_TEXT:-}" ]]; then
  log_error "Claude produced empty result"
  cat "$OUTPUT_FILE" >&2
  rm --force "$PROMPT_FILE" "$OUTPUT_FILE"
  exit 1
fi

log_info "Claude response received, extracting metadata..."

# Extract commit title (between ---COMMIT_TITLE--- and ---COMMIT_BODY---)
COMMIT_TITLE=$(echo "$RESULT_TEXT" | awk '/---COMMIT_TITLE---/{flag=1;next}/---COMMIT_BODY---/{flag=0}flag' | grep --invert-match '^[[:space:]]*$' | head --lines=1)

# Extract commit body (between ---COMMIT_BODY--- and ---PR_TITLE---)
COMMIT_BODY=$(echo "$RESULT_TEXT" | awk '/---COMMIT_BODY---/{flag=1;next}/---PR_TITLE---/{flag=0}flag')

# Extract PR title (between ---PR_TITLE--- and ---PR_DESCRIPTION---)
PR_TITLE=$(echo "$RESULT_TEXT" | awk '/---PR_TITLE---/{flag=1;next}/---PR_DESCRIPTION---/{flag=0}flag' | grep --invert-match '^[[:space:]]*$' | head --lines=1)

# Extract PR description (after ---PR_DESCRIPTION---)
PR_DESCRIPTION=$(echo "$RESULT_TEXT" | awk '/---PR_DESCRIPTION---/{flag=1;next}flag')

# Validate extractions
if [[ -z "${COMMIT_TITLE:-}" ]]; then
  log_error "Could not extract commit title from Claude output"
  echo "Claude output:" >&2
  echo "$RESULT_TEXT" >&2
  rm --force "$PROMPT_FILE" "$OUTPUT_FILE"
  exit 1
fi

if [[ -z "${PR_TITLE:-}" ]]; then
  log_error "Could not extract PR title from Claude output"
  echo "Claude output:" >&2
  echo "$RESULT_TEXT" >&2
  rm --force "$PROMPT_FILE" "$OUTPUT_FILE"
  exit 1
fi

# Output results using GitHub Actions output format with delimiters
echo "commit_title<<EOF_COMMIT_TITLE"
echo "$COMMIT_TITLE"
echo "EOF_COMMIT_TITLE"

echo "commit_body<<EOF_COMMIT_BODY"
echo "$COMMIT_BODY"
echo "EOF_COMMIT_BODY"

echo "pr_title<<EOF_PR_TITLE"
echo "$PR_TITLE"
echo "EOF_PR_TITLE"

echo "pr_description<<EOF_PR_DESCRIPTION"
echo "$PR_DESCRIPTION"
echo "EOF_PR_DESCRIPTION"

# Clean up
rm --force "$PROMPT_FILE" "$OUTPUT_FILE"

log_info "Metadata extracted successfully"
