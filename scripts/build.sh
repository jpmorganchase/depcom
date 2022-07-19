#!/bin/bash

rm -fr npm

# platform-windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-windows-64/bin/depcom.exe ./

# platform-windows-32
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -trimpath -o npm/depcom-windows-32/bin/depcom.exe ./

# platform-windows-arm64
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o npm/depcom-windows-arm64/bin/depcom.exe ./

# platform-darwin:
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-darwin-64/bin/depcom ./

# platform-darwin-arm64:
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o npm/depcom-darwin-arm64/bin/depcom ./

# platform-freebsd:
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-freebsd-64/bin/depcom ./

# platform-freebsd-arm64:
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o npm/depcom-freebsd-arm64/bin/depcom ./

# platform-netbsd:
CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-netbsd-64/bin/depcom ./

# platform-openbsd:
CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-openbsd-64/bin/depcom ./

# platform-linux:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-64/bin/depcom ./

# platform-linux-32:
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-32/bin/depcom ./

# platform-linux-arm:
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-arm/bin/depcom ./

# platform-linux-arm64:
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-arm64/bin/depcom ./

# platform-linux-mips64le:
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-mips64le/bin/depcom ./

# platform-linux-ppc64le:
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-ppc64le/bin/depcom ./

# platform-linux-riscv64:
CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-riscv64/bin/depcom ./

# platform-linux-s390x:
CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags="-s -w" -trimpath -o npm/depcom-linux-s390x/bin/depcom ./

# platform-sunos:
CGO_ENABLED=0 GOOS=illumos GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o npm/depcom-sunos-64/bin/depcom ./

