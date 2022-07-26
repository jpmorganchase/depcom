#!/bin/bash
set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/../npm
find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' && pwd && npm publish --registry http://localhost:4873/" \;
