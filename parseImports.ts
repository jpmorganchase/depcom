import fs from "fs-extra";
import esbuild from "esbuild";
import parseImports from "parse-imports";

async function getPackageDependencies(filePath: string) {
  const fileContents = fs.readFileSync(filePath, "utf-8");
  const { code } = esbuild.transformSync(fileContents, {
    target: "es2020",
    format: "esm",
    loader: "tsx", // Todo this varies depending on the file extension
  });
  const imports = [...(await parseImports(code))].filter(
    ({ moduleSpecifier }) => moduleSpecifier.type !== "relative"
  );
  console.log(JSON.stringify({ imports }, null, 2));
}

const cliArgs = process.argv.slice(2);
getPackageDependencies(cliArgs[0]);
