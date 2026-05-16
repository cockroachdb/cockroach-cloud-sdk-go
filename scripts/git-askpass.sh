#!/usr/bin/env bash
# GIT_ASKPASS script for git authentication. Reads credentials from
# environment variables so the token is never written to disk.
set -euo pipefail

# Get the script directory to find helper functions
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
source "$SCRIPT_DIR/lib/actions-helpers.sh"

if [ -z "$GIT_USER" ] || [ -z "$GIT_PASSWORD" ]; then
  log_error "GIT_USER and GIT_PASSWORD are required"
  exit 1
fi

case "$1" in
  Username*) echo "$GIT_USER" ;;
  Password*) echo "$GIT_PASSWORD" ;;
esac
