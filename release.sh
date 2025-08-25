#!/bin/bash
set -euo pipefail

# The most recent tag
echo "The most recent tag: $(git describe --tags --abbrev=0 2>/dev/null || echo 'none')"

# Prompt for version
read -p "Semantic version: " VERSION

# Validate version input
if [[ -z "$VERSION" ]]; then
    echo "Error: Version cannot be empty." >&2
    exit 1
elif [[ "$VERSION" == v* ]]; then
    echo "Error: Use the raw semantic version, without a 'v' prefix." >&2
    exit 1
elif ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in the format X.Y.Z (e.g., 1.2.3)." >&2
    exit 1
fi

# Get git commit hash
REV=$(git rev-parse --short HEAD)

# Create annotated tag
echo "Tagging $REV as v$VERSION"
git tag --annotate "v$VERSION" -m "Release v$VERSION"

# Push instructions
echo "Be sure to: git push --tags"
echo
