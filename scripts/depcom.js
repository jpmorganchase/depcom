const os = require("os");
const { execFile } = require("node:child_process");

const platform = os.platform();
const arch = os.arch();

const platformBinPath = require.resolve(
  `depcom-${platform()}-${arch()}/bin/depcom${arch === "win32" ? ".exe" : ""}`
);

function extractFromDirectory(directory) {
  return new Promise((resolve) => {
    execFile(platformBinPath, [`-d ${directory}`], (error, stdout) => {
      if (error) {
        throw error;
      }
      try {
        const result = JSON.parse(stdout);
        resolve(result);
      } catch (e) {
        throw new Error(`Can't parse depcom output:\n${stdout}`);
      }
    });
  });
}

function extractFromFiles(files) {}

module.exports = { extractFromDirectory, extractFromFiles };

console.log("Successfully resolved: ", platformBinPath);
