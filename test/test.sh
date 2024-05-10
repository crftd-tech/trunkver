#!/bin/bash

set -euo pipefail

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export TRUNKVER="$CURRENT_DIR/../trunkver.sh"
export REPO_DIR=$(mktemp -d)

{
    git init "$REPO_DIR"
    git -C "$REPO_DIR" config user.email "example@example.com"
    git -C "$REPO_DIR" config user.name "example@example.com"
    git -C "$REPO_DIR" commit --allow-empty -m "Initial commit"
} >/dev/null

export EXPECTED_GIT_SHA=$(git -C "$REPO_DIR" rev-parse --short HEAD)
export EXPECTED_TIME=$(date)

for test in "$CURRENT_DIR"/test_*.sh; do
    for ci in "$CURRENT_DIR/ci_setups"/*.sh; do
        set +e
        (
            date() {
                command date --date="${EXPECTED_TIME}" "$@"
            }
            export -f date
            . "$CURRENT_DIR/assert.sh"
            . "$ci"
            . "$test"
        )
        RETVAL=$?
        set -e
        if [ $RETVAL -ne 0 ]; then
            echo "✖ $(basename -s .sh "$test")/$(basename -s .sh "$ci") failed"
        else
            echo "✔ $(basename -s .sh "$test")/$(basename -s .sh "$ci") success"
        fi
    done
done
