package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ije/esbuild-internal/js_parser"
	"github.com/ije/esbuild-internal/logger"
	"github.com/ije/esbuild-internal/test"
)
func parseJsFile(filename string) () {	
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	log := logger.NewDeferLog(logger.DeferLogNoVerboseOrDebug)
	ast, pass := js_parser.Parse(log, test.SourceForTest(string(data)), js_parser.Options{})
	if pass {
		fmt.Println(ast)
	} else {
		fmt.Println(pass)
	}
	return
}

func main() {
	parseJsFile("../modular/packages/modular-scripts/react-scripts/scripts/start.js")
}