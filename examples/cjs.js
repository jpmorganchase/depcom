const foo = require("by-cjs-require");
let bar = require.resolve("by-require-resolve");
require("./local-import");
require("/absolute/import");

console.log(foo, bar);

console.log(`require${require("by-require-in-expression")}`);
