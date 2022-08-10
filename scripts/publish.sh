#!/bin/bash
set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/../npm

# Debug info
echo "Showing package.json for all the packages we are going to publish"
find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' && pwd && cat package.json" \;
echo "-----------------------------------------------------------------------------------------------------------------"
#

find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' && pwd && npm publish" \;
