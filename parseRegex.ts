import fs from "fs-extra";
import { globby } from "globby";
import strip from "strip-comments";

// https://github.com/jonschlinkert/requires-regex/blob/master/index.js
const requireRegex =
  /(?:(?:\\['"`][\s\S])*?(['"`](?=[\s\S]*?require\s*\(['"`][^`"']+?[`'"]\)))(?:\\\1|[\s\S])*?\1|\s*(?:(?:var|const|let)?\s*([_.\w/$]+?)\s*=\s*)?require\s*\(([`'"])((?:@([^/]+?)\/([^/]*?)|[-.@\w/$]+?))\3(?:, ([`'"])([^\7]+?)\7)?\);?)/g;

// https://gist.github.com/manekinekko/7e58a17bc62a9be47172
const importRegex = new RegExp(
  /import(?:["'\s]*([\w*${}\n\r\t, ]+)from\s*)?["'\s]["'\s](.*[@\w_-]+)["'\s].*;$/,
  "mg"
);
const dynamicImportRegex = new RegExp(
  /import\((?:["'\s]*([\w*{}\n\r\t, ]+)\s*)?["'\s](.*([@\w_-]+))["'\s].*\);$/,
  "mg"
);

async function getFileDependency(filePath: string) {
  const fileContents = strip(await fs.readFile(filePath, "utf-8"));
  const requires = fileContents.match(requireRegex);
  const imports = fileContents.match(importRegex);
  const dynamicImports = fileContents.match(dynamicImportRegex);
  // console.log(`Deps for file: ${filePath}`);
  // console.log(JSON.stringify({ imports }, null, 2));
  return { imports, requires, dynamicImports };
}

async function getPackageDependencies(files: string[]) {
  const dependenciesPerFile = {};
  const dependencies = {};
  for (const file of files) {
    const imports = await getFileDependency(file);
    //console.log("------------- FILE", file);
    //console.log(imports);
    // for (const importData of imports) {
    // if (!dependenciesPerFile[file]) {
    //   dependenciesPerFile[file] = [];
    // }
    // dependenciesPerFile[file].push(importData.moduleSpecifier.value);
    // if (!dependencies[importData.moduleSpecifier.value]) {
    //   dependencies[importData.moduleSpecifier.value] = { files: [] };
    // }
    // dependencies[importData.moduleSpecifier.value].files.push(file);
    // }
  }
  // return { dependenciesPerFile, dependencies };
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
