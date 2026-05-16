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

# Check if required environment variables are set
# Usage: check_required_env "VAR1" "VAR2" "VAR3"
check_required_env() {
  local missing_vars=()

  for var in "$@"; do
    if [[ -z "${!var:-}" ]]; then
      missing_vars+=("$var")
    fi
  done

  if [[ ${#missing_vars[@]} -gt 0 ]]; then
    log_error "Required environment variable(s): ${missing_vars[*]} are required"
    return 1
  fi

  return 0
}