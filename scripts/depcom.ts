import os from "os";
import { execFile } from "node:child_process";

interface AnalyzedDependencies {
  Time: string;
  ImportArray: string[];
  Logs: {
    Verbose: string[] | null;
    Debug: string[] | null;
    Info: string[] | null;
    Err: string[] | null;
    Warning: string[] | null;
    FileCount: number;
  };
}

interface AnalyzeRuntimeDependenciesParams {
  path: string;
  options: {
    match?: string;
    exclude?: string[];
  };
}

const platform = os.platform();
const arch = os.arch();

const platformBinPath = require.resolve(
  `depcom-${platform}-${arch}/bin/depcom${arch === "win32" ? ".exe" : ""}`
);

console.log("Successfully resolved: ", platformBinPath);

export function analyzeRuntimeDependencies({
  path,
  options: { match, exclude },
}: AnalyzeRuntimeDependenciesParams): Promise<AnalyzedDependencies> {
  return new Promise((resolve) => {
    execFile(platformBinPath, [`-d ${path}`], (error, stdout) => {
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
