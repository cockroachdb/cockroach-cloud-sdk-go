#!/bin/sh
# GIT_ASKPASS script for fork authentication. Reads credentials from
# environment variables so the token is never written to disk.
set -e

# Get the script directory to find helper functions
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
. "$SCRIPT_DIR/logging.sh"

if [ -z "$GIT_FORK_USER" ] || [ -z "$GIT_FORK_PASSWORD" ]; then
  log_error "GIT_FORK_USER and GIT_FORK_PASSWORD must be set"
  exit 1
fi

case "$1" in
  Username*) echo "$GIT_FORK_USER" ;;
  Password*) echo "$GIT_FORK_PASSWORD" ;;
esac
