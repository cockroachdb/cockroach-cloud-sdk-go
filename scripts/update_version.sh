#!/bin/bash

set -e

# This script is called by the reusable workflow after CHANGELOG.md has been updated.
# The VERSION environment variable is set by the reusable workflow.
# This script should only update version numbers in files other than CHANGELOG.md.

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source helper functions
source "$SCRIPT_DIR/lib/actions-helpers.sh"
source "$SCRIPT_DIR/lib/validation.sh"

# Check required commands are available
check_required_commands cut sed make

# Check if VERSION environment variable is set
if [ -z "$VERSION" ]; then
    log_error "VERSION environment variable is required"
    log_error "This script is intended to be called by the create-release-pr workflow"
    exit 1
fi

log_info "Updating version to $VERSION in config and documentation files..."

# Extract major version from VERSION (e.g., 7.0.0 -> 7)
MAJOR_VERSION=$(echo "$VERSION" | cut --delimiter=. --fields=1)

# Update go.mod module path to match major version
log_info "Updating go.mod..."
sed --regexp-extended "s|(github.com/cockroachdb/cockroach-cloud-sdk-go)/v[0-9]+|\1/v${MAJOR_VERSION}|" go.mod > go.mod.tmp && mv go.mod.tmp go.mod

# Update main.go import path to match major version
log_info "Updating main.go..."
sed --regexp-extended "s|(github.com/cockroachdb/cockroach-cloud-sdk-go)/v[0-9]+|\1/v${MAJOR_VERSION}|" main.go > main.go.tmp && mv main.go.tmp main.go

# Update internal/spec/config.yaml
log_info "Updating internal/spec/config.yaml..."
sed "s/^packageVersion:.*/packageVersion: $VERSION/" internal/spec/config.yaml > internal/spec/config.yaml.tmp && mv internal/spec/config.yaml.tmp internal/spec/config.yaml

# Generate OpenAPI client to update README files
log_info "Generating OpenAPI client..."
# Note: sudo is required because docker containers run as root by default,
# creating files owned by root. We then chown them back to the current user.
sudo make generate-openapi-client
sudo chown --recursive "$(id --user):$(id --group)" internal/openapi-generator/ pkg/client/ docs/ README.md

log_info "Version updated to $VERSION"

# Validate the updated code
log_info "Running validation..."
make validate