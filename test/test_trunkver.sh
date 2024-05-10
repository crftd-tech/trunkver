#!/bin/bash

set -euo pipefail

EXPECTED_VERSION=$(date --utc +%Y%m%d%H%M%S)
EXPECTED_JOB_ID=${EXPECTED_JOB_ID:-"123456789"}

assert_eq "$("$TRUNKVER")" "$EXPECTED_VERSION+g${EXPECTED_GIT_SHA}-${EXPECTED_JOB_ID}" "Version should be $EXPECTED_VERSION"
