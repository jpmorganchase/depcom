#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

rm -fr npm

# platform-windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-windows-64/bin/depcom.exe $SCRIPT_DIR/../

# platform-windows-32
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-windows-32/bin/depcom.exe $SCRIPT_DIR/../

# platform-windows-arm64
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-windows-arm64/bin/depcom.exe $SCRIPT_DIR/../

# platform-darwin:
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-darwin-64/bin/depcom $SCRIPT_DIR/../

# platform-darwin-arm64:
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-darwin-arm64/bin/depcom $SCRIPT_DIR/../

# platform-freebsd:
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-freebsd-64/bin/depcom $SCRIPT_DIR/../

# platform-freebsd-arm64:
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-freebsd-arm64/bin/depcom $SCRIPT_DIR/../

# platform-netbsd:
CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-netbsd-64/bin/depcom $SCRIPT_DIR/../

# platform-openbsd:
CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-openbsd-64/bin/depcom $SCRIPT_DIR/../

# platform-linux:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-64/bin/depcom $SCRIPT_DIR/../

# platform-linux-32:
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-32/bin/depcom $SCRIPT_DIR/../

# platform-linux-arm:
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-arm/bin/depcom $SCRIPT_DIR/../

# platform-linux-arm64:
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-arm64/bin/depcom $SCRIPT_DIR/../

# platform-linux-mips64le:
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-mips64le/bin/depcom $SCRIPT_DIR/../

# platform-linux-ppc64le:
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-ppc64le/bin/depcom $SCRIPT_DIR/../

# platform-linux-riscv64:
CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-riscv64/bin/depcom $SCRIPT_DIR/../

# platform-linux-s390x:
CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-s390x/bin/depcom $SCRIPT_DIR/../

# platform-sunos:
CGO_ENABLED=0 GOOS=illumos GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-sunos-64/bin/depcom $SCRIPT_DIR/..

node $SCRIPT_DIR/build.js
