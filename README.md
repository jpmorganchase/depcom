# depcom

A Go package that extracts imported dependencies from Javascript / Typescript / CSS files.
It uses [internals](https://github.com/ije/esbuild-internal/) from the [Esbuild project](https://esbuild.github.io/) to be blazing fast.

## Usage

### Build

`go build`

### Analyze a directory

`./depcom -d ../path/to/directory`

### Analyze single file

`./depcom -f ../path/to/directory/file.js`

### Analyze multiple files

`./depcom ../path/to/directory/file.js`

### Help

`./depcom -h`

## Supported import statements

- cjs [`require`](https://nodejs.org/api/modules.html#requireid) and [`require.resolve`](https://nodejs.org/api/modules.html#requireresolverequest-options), if the argument is a string literal.
- esm `import` [statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import) and [operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/import). The latter, commonly known as dynamic import, is supported only if the argument is a string literal.
- css [`@import` rule](https://developer.mozilla.org/en-US/docs/Web/API/CSSImportRule).

## Supported extension

- `.ts` - Typescript files.
- `.js` - Javascript files.
- `.jsx` - Javascript files with React JSX code. Please note that a file with extension `.js` containing JSX code will not be parsed correctly and will terminate parsing at the first JSX expression. This will emit an error in the logs but won't interrupt parsing of the remaining files.
- `.tsx` - Typescript files with React JSX code. Please note that a file with extension `.ts` containing JSX code will not be parsed correctly and will terminate parsing at the first JSX expression. This will emit an error in the logs but won't interrupt parsing of the remaining files.
- `.css` - CSS files

## Output

### Example

`json {"Time":"15.961751ms","ImportArray":["rollup-plugin-esbuild","jest-config","react-native-web","prettier","pptr-testing-library","tmp","address","detect-port-alt","is-ci","module","is-root","micromatch","rollup-plugin-postcss","modular-scripts","browserslist","update-notifier","semver-regex","util","filesize","@rollup/plugin-json","babel-jest","prompts","express-ws","recursive-readdir","@rollup/plugin-node-resolve","escape-string-regexp","tree-view-for-tests","dotenv","esbuild","react-error-overlay","child_process","cross-spawn","dedent","ts-morph","@schemastore/tsconfig","express","jest-circus","url","fs-extra","source-map-support","ts-jest","find-up","mime","react-dom","chalk","gzip-size","npm-packlist","commander","globby","stream","@babel/code-frame","@rollup/plugin-commonjs","typescript","rimraf","dotenv-expand","http","foo","builtin-modules","rollup","babel-preset-react-app","html-minifier-terser","execa","puppeteer","semver","resolve","os","ws","change-case","jest-cli","jest-transform-stub","js-yaml","open","parse5","path","jest-watch-typeahead","strip-ansi","@svgr/core","react"],"Logs":{"Verbose":null,"Debug":["../modular/packages/modular-scripts/src/check/index.ts: This \"import\" expression will not be bundled because the argument is not a string literal\n","../modular/packages/modular-scripts/src/esbuild-scripts/start/index.ts: This call to \"require\" will not be bundled because the argument is not a string literal\n"],"Info":null,"Err":null,"Warning":null},"FileCount":119}`

### Output format

The `ImportArray` field an array of unique dependencies extracted from the files specified in the arguments, without any indication regarding how or in which file the dependency is imported or the [subpath](https://nodejs.org/api/packages.html#subpath-patterns) imported.
