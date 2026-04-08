#!/bin/bash

set -e

# This script is called by the reusable workflow after CHANGELOG.md has been updated.
# The VERSION environment variable is set by the reusable workflow.
# This script should only update version numbers in files other than CHANGELOG.md.

# Check if VERSION environment variable is set
if [ -z "$VERSION" ]; then
    echo "Error: VERSION environment variable must be set"
    echo "This script is intended to be called by the create-release-pr workflow"
    exit 1
fi

echo "Updating version to $VERSION in config and documentation files..."

# Extract major version from VERSION (e.g., 7.0.0 -> 7)
MAJOR_VERSION=$(echo "$VERSION" | cut -d. -f1)

# Update go.mod module path to match major version
echo "Updating go.mod..."
sed -E "s|(github.com/cockroachdb/cockroach-cloud-sdk-go)/v[0-9]+|\1/v${MAJOR_VERSION}|" go.mod > go.mod.tmp && mv go.mod.tmp go.mod

# Update internal/spec/config.yaml
echo "Updating internal/spec/config.yaml..."
sed "s/^packageVersion:.*/packageVersion: $VERSION/" internal/spec/config.yaml > internal/spec/config.yaml.tmp && mv internal/spec/config.yaml.tmp internal/spec/config.yaml

# Update README.md
echo "Updating README.md..."
sed "s/^- Package version:.*/- Package version: $VERSION/" README.md > README.md.tmp && mv README.md.tmp README.md

# Update docs/README.md
echo "Updating docs/README.md..."
sed "s/^- Package version:.*/- Package version: $VERSION/" docs/README.md > docs/README.md.tmp && mv docs/README.md.tmp docs/README.md

# Update pkg/client/configuration.go UserAgent
echo "Updating pkg/client/configuration.go..."
sed "s/UserAgent:     \"ccloud-sdk-go\/[^\"]*\"/UserAgent:     \"ccloud-sdk-go\/$VERSION\"/" pkg/client/configuration.go > pkg/client/configuration.go.tmp && mv pkg/client/configuration.go.tmp pkg/client/configuration.go

echo "Version updated to $VERSION"
