import esbuild from "esbuild";
import { globby } from "globby";

async function getFileDependency(filePath: string) {
  let detectedImports: esbuild.OnResolveArgs[] = [];
  await esbuild.build({
    target: "es2020",
    format: "esm",
    entryPoints: [filePath],
    write: false,
    outfile: "out.js",
    bundle: true,
    plugins: [
      {
        name: "detectImportsPlugin",
        setup: (build) => {
          build.onResolve({ filter: /.*/ }, (args) => {
            // console.log(`Found an import: ${JSON.stringify(args)}`);
            if (args.kind === "entry-point") {
              return {};
            }
            if (!isImportLocal(args.path)) {
              detectedImports.push(args);
            }
            return { external: true };
          });
        },
      },
    ],
  });
  return detectedImports;
}

function isImportLocal(importPath: string) {
  return importPath.startsWith(".");
}

async function getPackageDependencies(files: string[]) {
  const dependenciesPerFile = {};
  const dependencies = {};
  for (const file of files) {
    const imports = await getFileDependency(file);
    for (const importArgs of imports) {
      if (!dependenciesPerFile[file]) {
        dependenciesPerFile[file] = [];
      }
      dependenciesPerFile[file].push(importArgs.path);
      if (!dependencies[importArgs.path]) {
        dependencies[importArgs.path] = { files: [] };
      }
      dependencies[importArgs.path].files.push(file);
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

//console.log(JSON.stringify(result.dependencies, null, 2));
