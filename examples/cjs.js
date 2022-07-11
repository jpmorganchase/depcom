const lodash = require("lodash");
var react = require("react");
let jj = require.resolve("jj");
require("./my_stuff");

console.log(react, lodash, jj);

console.log(`require${require("globby")}`);
