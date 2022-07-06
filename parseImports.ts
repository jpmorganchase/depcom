import fs from "fs-extra";
import esbuild from "esbuild";
import parseImports from "parse-imports";
import { globby } from "globby";

async function getFileDependency(filePath: string) {
  const fileContents = await fs.readFile(filePath, "utf-8");
  const { code } = await esbuild.transform(fileContents, {
    target: "es2020",
    format: "esm",
    loader: "tsx", // Todo this varies depending on the file extension
  });
  const imports = [...(await parseImports(code))].filter(
    ({ moduleSpecifier }) => moduleSpecifier.type !== "relative"
  );
  // console.log(`Deps for file: ${filePath}`);
  // console.log(JSON.stringify({ imports }, null, 2));
  return imports;
}

async function getPackageDependencies(files: string[]) {
  const dependenciesPerFile = {};
  const dependencies = {};
  for (const file of files) {
    const imports = await getFileDependency(file);
    for (const importData of imports) {
      if (!dependenciesPerFile[file]) {
        dependenciesPerFile[file] = [];
      }
      dependenciesPerFile[file].push(importData.moduleSpecifier.value);
      if (!dependencies[importData.moduleSpecifier.value]) {
        dependencies[importData.moduleSpecifier.value] = { files: [] };
      }
      dependencies[importData.moduleSpecifier.value].files.push(file);
    }
  }
  return { dependenciesPerFile, dependencies };
}

export async function getDependencies(directoryPath: string) {
  return getPackageDependencies(
    await globby(directoryPath, {
      expandDirectories: {
        extensions: ["ts", "tsx", "js", "jsx"],
      },
    })
  );
}

const cliArgs = process.argv.slice(2);

console.time("getDependencies");
const result = await getDependencies(cliArgs[0]);
console.timeEnd("getDependencies");

// console.log(JSON.stringify(result.dependencies, null, 2));
