#!/bin/bash

# =================================================================
#
# Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

testOne() {
  local expected='hello planet'
  local output=$(echo 'hello world' | goreplacer -r '[["world", "planet"]]')
  assertEquals "unexpected output" "$(echo -e "${expected}")" "${output}"
}

testDoubleQuotes() {
  local expected='{"hello":"world"}'
  local output=$(echo '{""hello"":""world""}' | goreplacer -r '[["\"\"","\""]]')
  assertEquals "unexpected output" "$(echo -e "${expected}")" "${output}"
}

testSingleQuotesToDoubleQuotes() {
  local expected='{"hello":"world"}'
  local output=$(echo "{'hello':'world'}" | goreplacer -r "[[\"'\",\"\\\"\"]]")
  assertEquals "unexpected output" "$(echo -e "${expected}")" "${output}"
}

oneTimeSetUp() {
  echo "Setting up"
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"
