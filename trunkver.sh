#!/bin/bash

set -euo pipefail

TIMESTAMP=$(date --utc +%Y%m%d%H%M%S)
GIT_REF=$(git -C "$GITHUB_WORKSPACE" rev-parse --short HEAD)
BUILD_ID=${GITHUB_RUN_ID}_${GITHUB_RUN_ATTEMPT}

echo "$TIMESTAMP+g$GIT_REF-$BUILD_ID"