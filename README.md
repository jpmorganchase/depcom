# depcom

A Go package that extracts imported dependencies from Javascript / Typescript / CSS source files.
It uses heavy parallelization and [internal APIs](https://github.com/ije/esbuild-internal/) from the [Esbuild project](https://esbuild.github.io/) for blazing performance.

## Usage

### Build

`go build`

### Analyze a directory

`./depcom -d ../path/to/directory`

The `"/**/*.{tsx,jsx,mjs,cjs,ts,js,css}"` glob pattern will be appended to the specified directory.

### Analyze multiple files

`./depcom ../path/to/directory/file1.js ../another/path/to/directory/file1.js`

### Help

`./depcom -h`

## Supported import statements

- CJS [`require`](https://nodejs.org/api/modules.html#requireid) and [`require.resolve`](https://nodejs.org/api/modules.html#requireresolverequest-options), if the argument is a string literal.
- ESM `import` [statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import) and [operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/import). The latter, commonly known as dynamic import, is supported only if the argument is a string literal.
- CSS [`@import` rule](https://developer.mozilla.org/en-US/docs/Web/API/CSSImportRule).

## Supported file extensions

- `.ts` - Typescript files.
- `.js` - Javascript files.
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
