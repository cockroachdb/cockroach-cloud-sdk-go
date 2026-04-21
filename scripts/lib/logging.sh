#!/usr/bin/env bash
# Logging functions for GitHub Actions workflows

# Output an informational message to stdout
log_info() {
  local message="$1"
  echo "$message"
}

# Output an error message using GitHub Actions workflow command format
# https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-error-message
log_error() {
  local message="$1"
  echo "::error::$message" >&2
}

# Output a warning message using GitHub Actions workflow command format
log_warning() {
  local message="$1"
  echo "::warning::$message" >&2
}

# Output a notice message using GitHub Actions workflow command format
log_notice() {
  local message="$1"
  echo "::notice::$message"
}
