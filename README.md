# depcom

A Go package that extracts imported dependencies from Javascript / Typescript / CSS source files.
It uses heavy parallelization and [internal APIs](https://github.com/ije/esbuild-internal/) from the [Esbuild project](https://esbuild.github.io/) for blazing performance.

## NPM Package

### Installation

`npm install depcom --save`

or

`yarn add depcom`

### Usage

```ts
import { analyzeRuntimeDependencies } from "depcom";

const { ImportArray, Time, FileCount, Logs } = analyzeRuntimeDependencies({
  path: "path/to/package",
  options: {
    match: "**/*.{tsx,jsx,mjs,cjs,ts,js,css}",
    exclude: ["node_modules/**/*"],
  },
});
```

## CLI

### Usage

#### Build

`go build`

#### Tests

`go test ./...`

#### CLI Options

##### Match files

- `-d` Set a base directory (default: `./`)
- `-a` Select multiple files using a [glob pattern](https://github.com/mattn/go-zglob), starting from the base directory (default: `**/*.{tsx,jsx,mjs,cjs,ts,js,css}`)
- `-x` Exclude files using a [glob pattern](https://github.com/mattn/go-zglob), starting from the base directory. This option can be specified multiple times (default: none)

Target files will be matched by evaluating the glob patterns separately, then calculating the difference between the allowed matches and all the excluded ones.

###### Examples

- Parse all javascript files in a package, excluding the `node_modules` directory (note the quotes, to avoid shell globbing):

`./depcom -d path/to/package -a "**/*.{tsx,jsx,mjs,cjs,ts,js,css}" -x "node_modules/**/*"`

- Parse all javascript files in a package outside of `src` that aren't external dependencies (note the double usage of the -x argument):

`./depcom -d path/to/package -a "**/*.{tsx,jsx,mjs,cjs,ts,js,css}" -x "node_modules/**/*" -x "src/**/*"`

- Parse all the javascript files in the current directory and subdirectories, recursively

`./depcom`

##### Variadic usage

`./depcom ../path/to/directory/file1.js ../another/path/to/directory/file1.js`

#### Help

`./depcom -h`

## Supported import statements

- CJS [`require`](https://nodejs.org/api/modules.html#requireid) and [`require.resolve`](https://nodejs.org/api/modules.html#requireresolverequest-options), if the argument is a string literal.
- ESM `import` [statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import) and [operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/import). The latter, commonly known as dynamic import, is supported only if the argument is a string literal.
- CSS [`@import` rule](https://developer.mozilla.org/en-US/docs/Web/API/CSSImportRule).

## Supported file extensions

- `.js` - Javascript files. All unrecognized extensions will fall back to `.js` (so, for example, you can safely pass `.mjs` or `.cjs` files to depcom)
- `.ts` - Typescript files.
- `.jsx` - Javascript files with React JSX code. Please note that a file with extension `.js` containing JSX code will not be parsed correctly and will terminate parsing at the first JSX expression. This will emit an error in the logs but won't interrupt parsing of the remaining files.
- `.tsx` - Typescript files with React JSX code. Please note that a file with extension `.ts` containing JSX code will not be parsed correctly and will terminate parsing at the first JSX expression. This will emit an error in the logs but won't interrupt parsing of the remaining files.
- `.css` - CSS files

## Output

### Format

- `Time` - Time elapsed parsing
- `Logs` - Array of logs, grouped by log level
- `ImportArray` - An array of all the unique dependencies extracted from the files. No subpaths.
- `FileCount` - The number of files processed

### Example

`json {"Time":"15.961751ms","ImportArray":["rollup-plugin-esbuild","jest-config","react-native-web",...],"Logs":{"Verbose":null,"Debug":["../modular/packages/modular-scripts/src/check/index.ts: This \"import\" expression will not be bundled because the argument is not a string literal\n","../modular/packages/modular-scripts/src/esbuild-scripts/start/index.ts: This call to \"require\" will not be bundled because the argument is not a string literal\n"],"Info":null,"Err":null,"Warning":null},"FileCount":119}`
