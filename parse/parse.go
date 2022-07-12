package parse

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
)

func FromFiles(filenames []string) []string {
	outChannel := make(chan []string, len(filenames))
	importSet := make(map[string]bool)
	var result []string

	for _, match := range filenames {
		go fromFileAsync(match, outChannel)
	}

	for i := 0; i < len(filenames); i++ {
		fileImports := <-outChannel
		for _, fileImport := range fileImports {
			importSet[fileImport] = true
		}
	}

	for k := range importSet {
		result = append(result, k)
	}
	return result
}

func fromFileAsync(filename string, outChannel chan []string) {
	outChannel <- FromFile(filename)
}

func FromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("** Error reading file: %s - %#v\n", filename, err)
		return make([]string, 0)
	}
	log := logger.NewStderrLog(logger.OutputOptions{LogLevel: logger.LevelDebug})
	ext := filepath.Ext(filename)

	// fmt.Printf("-----> PARSING: [%v ext: %v]\n", filename, ext)

	if ext == ".css" {
		return fromCSS(&log, &data)
	} else {
		return fromECMA(&log, &data, ext)
	}

}

func fromCSS(log *logger.Log, data *[]byte) []string {
	var imports []string

	ast := css_parser.Parse(*log, test.SourceForTest(string(*data)), css_parser.Options{})

	for _, record := range ast.ImportRecords {
		if !isDependencyLocal(record.Path.Text) {
			// fmt.Printf("%#v\n", record.Path.Text)
			imports = append(imports, record.Path.Text)
		}
	}
	return imports
}

func fromECMA(log *logger.Log, data *[]byte, ext string) []string {
	var imports []string

	options := config.Options{Mode: 2}
	if ext == ".ts" || ext == ".tsx" {
		// fmt.Printf("setting ts for %v\n", ext)
		options.TS = config.TSOptions{
			Parse: true,
		}
	}

	if ext == ".jsx" || ext == ".tsx" {
		options.JSX = config.JSXOptions{
			Parse: true,
		}
	}

	ast, pass := js_parser.Parse(*log, test.SourceForTest(string(*data)), js_parser.OptionsFromConfig(&options))

	if pass {
		for _, record := range ast.ImportRecords {
			if !isDependencyLocal(record.Path.Text) {
				// fmt.Printf("%#v\n", record.Path.Text)
				imports = append(imports, record.Path.Text)
			}
		}
	} else {
		fmt.Println("** Not passed!")
	}

	return imports
}

func isDependencyLocal(dependency string) bool {
	return (strings.HasPrefix(dependency, ".") || strings.HasPrefix(dependency, "/"))
}
