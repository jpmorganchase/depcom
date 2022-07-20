const { execSync } = require("node:child_process");
const fs = require("fs");
const path = require("path");

const archMap = {
  "depcom-windows-arm64": {
    os: "win32",
    cpu: "arm64",
    endianness: "LE",
    goos: "windows",
    goarch: "arm64",
    extension: ".exe",
  },
  "depcom-windows-32": {
    os: "win32",
    cpu: "ia32",
    goos: "windows",
    goarch: "386",
    endianness: "LE",
    extension: ".exe",
  },
  "depcom-windows-64": {
    os: "win32",
    cpu: "x64",
    goos: "windows",
    goarch: "amd64",
    endianness: "LE",
    extension: ".exe",
  },
  "depcom-darwin-arm64": {
    os: "darwin",
    cpu: "arm64",
    goos: "darwin",
    goarch: "arm64",
    endianness: "LE",
  },
  "depcom-darwin-64": {
    os: "darwin",
    cpu: "x64",
    goos: "darwin",
    goarch: "amd64",
    endianness: "LE",
  },
  "depcom-freebsd-arm64": {
    os: "freebsd",
    cpu: "arm64",
    goos: "freebsd",
    goarch: "arm64",
    endianness: "LE",
  },
  "depcom-freebsd-64": {
    os: "freebsd",
    cpu: "x64",
    goos: "freebsd",
    goarch: "amd64",
    endianness: "LE",
  },
  "depcom-linux-arm": {
    os: "linux",
    cpu: "arm",
    goos: "linux",
    goarch: "arm",
    endianness: "LE",
  },
  "depcom-linux-arm64": {
    os: "linux",
    cpu: "arm64",
    goos: "linux",
    goarch: "arm64",
    endianness: "LE",
  },
  "depcom-linux-32": {
    os: "linux",
    cpu: "ia32",
    goos: "linux",
    goarch: "386",
    endianness: "LE",
  },
  "depcom-linux-mips64le": {
    os: "linux",
    cpu: "mips64el",
    goos: "linux",
    goarch: "mips64le",
    endianness: "LE",
  },
  "depcom-linux-ppc64le": {
    os: "linux",
    cpu: "ppc64",
    goos: "linux",
    goarch: "ppc64le",
    endianness: "LE",
  },
  "depcom-linux-riscv64": {
    os: "linux",
    cpu: "riscv64",
    goos: "linux",
    goarch: "riscv64",
    endianness: "LE",
  },
  "depcom-linux-s390x": {
    os: "linux",
    cpu: "s390x",
    goos: "linux",
    goarch: "s390x",
    endianness: "BE",
  },
  "depcom-linux-64": {
    os: "linux",
    cpu: "x64",
    goos: "linux",
    goarch: "amd64",
    endianness: "LE",
  },
  "depcom-netbsd-64": {
    os: "netbsd",
    cpu: "x64",
    goos: "netbsd",
    goarch: "amd64",
    endianness: "LE",
  },
  "depcom-openbsd-64": {
    os: "openbsd",
    cpu: "x64",
    goos: "openbsd",
    goarch: "amd64",
    endianness: "LE",
  },
  "depcom-sunos-64": {
    os: "sunos",
    cpu: "x64",
    goos: "illumos",
    goarch: "amd64",
    endianness: "LE",
  },
};

function main() {
  Object.entries(archMap).map(
    ([target, { os, cpu /*, extension, goos, goarch */ }]) => {
      const baseManifest = require(path.join(__dirname, "../package.json"));
      // build(target, goos, goarch, extension);
      generateManifest(target, os, cpu, baseManifest);
    }
  );
}

function generateManifest(target, os, cpu, baseManifest) {
  const newManifest = {
    name: target,
    version: baseManifest.version,
    description: `depcom - ${target} build`,
    repository: baseManifest.repository,
    license: baseManifest.license,
    preferUnplugged: false,
    engines: baseManifest.engines,
    os: [os],
    cpu: [cpu],
  };

  fs.writeFileSync(
    path.join(__dirname, "../npm/", target, "package.json"),
    JSON.stringify(newManifest, null, 2)
  );
}

function build(target, goos, goarch, extension = "") {
  // TODO: this doesn't work atm, requires $GOPATH to be set
  const rootDirectory = path.join(__dirname, "../");
  const command = `go build -ldflags="-s -w" -trimpath -o ./npm/${target}/bin/depcom${extension} ./`;

  const res = execSync(command, {
    env: {
      CGO_ENABLED: 0,
      GOOS: goos,
      GOARCH: goarch,
    },
    cwd: rootDirectory,
  });
}

module.exports = {
  archMap,
};

main();
