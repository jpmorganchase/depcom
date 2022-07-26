#!/bin/bash
set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/../npm

# TODO debug info, remove
find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' && pwd && cat package.json" \;
#

find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' && pwd && npm publish" \;
