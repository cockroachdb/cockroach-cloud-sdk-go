#!/usr/bin/env bash
# Validation helper functions

# Check if required commands are available in PATH
# Usage: check_required_commands "cmd1" "cmd2" "cmd3"
check_required_commands() {
  local missing_commands=()

  for cmd in "$@"; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
      missing_commands+=("$cmd")
    fi
  done

  if [[ ${#missing_commands[@]} -gt 0 ]]; then
    log_error "Required command(s) not found in PATH: ${missing_commands[*]}"
    return 1
  fi

  return 0
}
