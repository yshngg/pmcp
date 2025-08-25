#!/bin/bash

set -euo pipefail

# The most recent tag, or 'none' if no tags exist
RECENT=$(git describe --tags --abbrev=0 2>/dev/null || echo 'none')
echo "The most recent tag: $RECENT"

# Show commit logs since the most recent tag
echo "The commit logs since the most recent tag:"
if [[ "$RECENT" == 'none' ]]; then
    git log --oneline
else
    if [[ -z $(git rev-list --max-count=1 "${RECENT}..HEAD") ]]; then
        echo "No new commits since tag $RECENT. Aborting."
        exit 1
    fi
    git log --oneline "${RECENT}...HEAD"
fi

# Prompt user for the new semantic version (e.g., 1.2.3)
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

# Check if tag already exists
if git rev-parse "v$VERSION" >/dev/null 2>&1; then
    echo "Error: Tag v$VERSION already exists. To delete, run: git tag -d v$VERSION" >&2
    exit 1
fi

# The current git commit short hash
REV=$(git rev-parse --short HEAD)

# Create an annotated tag for the release
echo "Tagging commit $REV as v$VERSION"
git tag --annotate "v$VERSION" -m "Release v$VERSION"

# Remind user to push tags to remote
echo "Release tag created. To publish, run: git push --tags"
