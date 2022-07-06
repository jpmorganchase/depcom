import fs from "fs-extra";
import parser from "@babel/parser";
import { globby } from "globby";

async function getFileDependency(filePath: string) {
  const fileContents = await fs.readFile(filePath, "utf-8");

  const parsed = parser.parse(fileContents, {
    sourceType: "unambiguous",
    plugins: [
      // enable jsx syntax
      "jsx",
      "typescript",
    ],
  });

  // console.log(parsed);

  return parsed;
}

async function getPackageDependencies(files: string[]) {
  const dependenciesPerFile = {};
  const dependencies = {};
  for (const file of files) {
    const imports = await getFileDependency(file);
    // for (const importData of imports) {
    //   if (!dependenciesPerFile[file]) {
    //     dependenciesPerFile[file] = [];
    //   }
    //   dependenciesPerFile[file].push(importData.moduleSpecifier.value);
    //   if (!dependencies[importData.moduleSpecifier.value]) {
    //     dependencies[importData.moduleSpecifier.value] = { files: [] };
    //   }
    //   dependencies[importData.moduleSpecifier.value].files.push(file);
    // }
  }
  return {};
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
