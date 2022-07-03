import esbuild from "esbuild";

async function getPackageDependencies(filePath: string) {
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
            console.log(`Found an import: ${JSON.stringify(args)}`);
            if (isImportLocal(args.path)) {
              return {};
            } else {
              detectedImports.push(args);
              // TODO if we detect it's in workspace, mark it as a workspace dependency too
              return { external: true };
            }
          });
        },
      },
    ],
  });
  console.log(detectedImports);
}

function isImportLocal(importPath: string) {
  return importPath.startsWith(".");
}

const cliArgs = process.argv.slice(2);
getPackageDependencies(cliArgs[0]);
