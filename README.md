# depcom
## esbuild way
- Scan all packages
- Build each one with esbuild. Use a plugin which country dependencies 
- Save each package's src/ hash + deps
- If some src/ change, recalculate
- Ignore external and local deps, look only for workspace deps

### pros:
- calculates the real deps (orphaned files are not taken into account)
- works with any type of file or module

### cons
- will need plugins (for example, workers?)
- package granularity (not file)

## parse-imports way

- Still use esbuild to translate in memory to esm
- Use https://github.com/TomerAberbach/parse-imports
- Save each file hash
- If some file in src changes, recalculate

### pros
- doesn't technically need plugins
- file granularity 

## cons
- not sure if it works with cjs modules
- will report orphaned files (but they can be excluded with configuration?)

## TODO
- [ ] poc of esbuild (just full scan)
- [ ] poc of parse-imports (just full scan)
- [ ] benchmark
