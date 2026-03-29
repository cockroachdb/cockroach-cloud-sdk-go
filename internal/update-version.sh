#!/bin/bash

set -e

# Check if version argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 6.11.0"
    exit 1
fi

VERSION="$1"
DATE=$(date +%Y-%m-%d)

# Validate semver format
if ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in semver format (e.g., 6.11.0)"
    exit 1
fi

# Helper function to set output
set_output() {
    echo "has_unreleased=$1"
    [ -n "$GITHUB_OUTPUT" ] && echo "has_unreleased=$1" >> "$GITHUB_OUTPUT"
}

# Check if [Unreleased] section exists
echo "Checking CHANGELOG.md for Unreleased section..."
if ! grep -q "^## \[Unreleased\]" CHANGELOG.md; then
    echo "No Unreleased section found. No files updated."
    set_output false
    exit 0
fi

# Check if [Unreleased] section has content
echo "Checking if Unreleased section has content..."
unreleased_content=$(awk '
/^## \[Unreleased\]/ { flag=1; next }
/^## \[/ { flag=0 }
flag && /^[^[:space:]]/ { print }
' CHANGELOG.md | grep -v "^$")

if [ -z "$unreleased_content" ]; then
    echo "Unreleased section is empty. No files updated."
    set_output false
    exit 0
fi

echo "Found Unreleased section. Updating version to $VERSION (date: $DATE)"

# Update CHANGELOG.md - keep [Unreleased] and add new version below
echo "Updating CHANGELOG.md..."
awk -v version="$VERSION" -v date="$DATE" '
/^## \[Unreleased\]/ { print; print ""; print "## [" version "] - " date; next }
{ print }
' CHANGELOG.md > CHANGELOG.md.tmp && mv CHANGELOG.md.tmp CHANGELOG.md

# Update internal/spec/config.yaml
echo "Updating internal/spec/config.yaml..."
sed "s/^packageVersion:.*/packageVersion: $VERSION/" internal/spec/config.yaml > internal/spec/config.yaml.tmp && mv internal/spec/config.yaml.tmp internal/spec/config.yaml

# Update README.md
echo "Updating README.md..."
sed "s/^- Package version:.*/- Package version: $VERSION/" README.md > README.md.tmp && mv README.md.tmp README.md

echo ""
echo "✓ Version updated to $VERSION in:"
echo "  - CHANGELOG.md"
echo "  - internal/spec/config.yaml"
echo "  - README.md"
echo ""

set_output true
