package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ije/esbuild-internal/config"
	"github.com/ije/esbuild-internal/css_parser"
	"github.com/ije/esbuild-internal/js_parser"
	"github.com/ije/esbuild-internal/logger"
	"github.com/ije/esbuild-internal/test"
	"github.com/mattn/go-zglob"
)

func isDependencyLocal(dependency string) bool {
	return (strings.HasPrefix(dependency, ".") || strings.HasPrefix(dependency, "/"))
}

func parseFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s - %#v\n", filename, err)
		return
	}
	log := logger.NewStderrLog(logger.OutputOptions{LogLevel: logger.LevelDebug})

	options := config.Options{}
	ext := filepath.Ext(filename)

	if ext == ".css" {
		ast := css_parser.Parse(log, test.SourceForTest(string(data)), css_parser.Options{})

		// fmt.Printf("%#v\n", ast)
		for _, record := range ast.ImportRecords {
			if !isDependencyLocal(record.Path.Text) {
				fmt.Printf("%#v\n", record.Path.Text)
			}
		}

		return
	}

	if ext == ".ts" || ext == ".tsx" {
		fmt.Printf("setting ts for %v\n", ext)
		options.TS = config.TSOptions{
			Parse: true,
		}
	}

	if ext == ".jsx" || ext == ".tsx" {
		options.JSX = config.JSXOptions{
			Parse: true,
		}
	}

	fmt.Printf("-----> PARSING: [%v ext: %v]\n", filename, ext)

	ast, pass := js_parser.Parse(log, test.SourceForTest(string(data)), js_parser.OptionsFromConfig(&options))

	if pass {
		for _, record := range ast.ImportRecords {
			if !isDependencyLocal(record.Path.Text) {
				fmt.Printf("%#v\n", record.Path.Text)
			}
		}
	} else {
		fmt.Println("Not passed!")
	}
}

func globMatches(dirPath string) {
	// TODO: exclude node_modules if possible (maybe manually?)
	matches, err := zglob.Glob(dirPath + "/**/*.{tsx,jsx,ts,js,css}")

	if err != nil {
		fmt.Println(err)
	} else {
		for _, match := range matches {
			parseFile(match)
		}
	}
}

func main() {
	globMatches("./examples")
}
