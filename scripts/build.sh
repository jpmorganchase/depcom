#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

rm -fr npm

echo "Building depcom-win32-arm64"
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-win32-arm64/bin/depcom.exe $SCRIPT_DIR/../

echo "Building depcom-win32-ia32"
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-win32-ia32/bin/depcom.exe $SCRIPT_DIR/../

echo "Building depcom-win32-x64"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-win32-x64/bin/depcom.exe $SCRIPT_DIR/../

echo "Building depcom-darwin-arm64"
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-darwin-arm64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-darwin-x64"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-darwin-x64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-freebsd-arm64"
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-freebsd-arm64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-freebsd-x64"
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-freebsd-x64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-arm"
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-arm/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-arm64"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-arm64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-ia32"
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-ia32/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-mips64el"
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-mips64el/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-ppc64"
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-ppc64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-riscv64"
CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-riscv64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-s390x"
CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-s390x/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-linux-x64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-linux-x64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-netbsd-x64"
CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-netbsd-x64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-openbsd-x64"
CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-openbsd-x64/bin/depcom $SCRIPT_DIR/../

echo "Building depcom-sunos-x64"
CGO_ENABLED=0 GOOS=illumos GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $SCRIPT_DIR/../npm/depcom-sunos-x64/bin/depcom $SCRIPT_DIR/../

echo "Generating package manifests"
node $SCRIPT_DIR/generate-manifests.js

cp $SCRIPT_DIR/depcom.js $SCRIPT_DIR/../npm/depcom/

