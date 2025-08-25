#!/bin/bash

# The most recent tag
echo "The most recent tag: $(git describe --tags --abbrev=0)"

# Prompt for version
read -p "Semantic version: " VERSION

# Check if version starts with 'v'
if [[ $VERSION == v* ]]; then
    echo "Use the raw semantic version, without a v prefix"
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
