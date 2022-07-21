const { execSync } = require("node:child_process");
const fs = require("fs");
const path = require("path");
const archMap = require("./architecture-map.json");

const baseManifest = require(path.join(__dirname, "../package.json"));

function main() {
  archMap.map(({ platform, arch, extension }) => {
    const target = `depcom-${platform}-${arch}`;
    generateManifest(target, platform, arch, extension);
    generateMainManifest();
  });
}

function generateManifest(target, os, cpu, extension = "") {
  const newManifest = {
    name: target,
    version: baseManifest.version,
    description: `depcom - ${target} build`,
    repository: baseManifest.repository,
    license: baseManifest.license,
    preferUnplugged: false,
    engines: baseManifest.engines,
    bin: `bin/depcom${extension}`,
    os: [os],
    cpu: [cpu],
  };

  fs.writeFileSync(
    path.join(__dirname, "../npm/", target, "package.json"),
    JSON.stringify(newManifest, null, 2)
  );
}

function generateMainManifest() {
  const mainPath = path.join(__dirname, "../npm/depcom/");
  fs.mkdirSync(mainPath, { recursive: true });
  const optionalDependencies = {};
  archMap.forEach(({ platform, arch }) => {
    optionalDependencies[`depcom-${platform}-${arch}`] = baseManifest.version;
  });
  fs.writeFileSync(
    path.join(mainPath, "package.json"),
    JSON.stringify({ ...baseManifest, optionalDependencies }, null, 2)
  );
}

function build(target, goos, goarch, extension = "") {
  // TODO: this doesn't work atm, requires $GOPATH to be set
  const rootDirectory = path.join(__dirname, "../");
  const command = `go build -ldflags="-s -w" -trimpath -o ./npm/${target}/bin/depcom${extension} ./`;

  execSync(command, {
    env: {
      CGO_ENABLED: 0,
      GOOS: goos,
      GOARCH: goarch,
    },
    cwd: rootDirectory,
  });
}

main();
