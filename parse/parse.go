package parse

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ije/esbuild-internal/config"
	"github.com/ije/esbuild-internal/css_parser"
	"github.com/ije/esbuild-internal/js_parser"
	"github.com/ije/esbuild-internal/logger"
)

type Imports struct {
	ImportArray []string
	Logs        LogMap
	FileCount   int
}

func FromFiles(filenames []string) *Imports {
	nameRegex, _ := regexp.Compile(`^(@[a-z0-9-~][a-z0-9-._~]*)?\/?([a-z0-9-~][a-z0-9-._~]*)`)
	maxGoroutines := 16
	// Throttle goroutines to maxGoroutines simultaneously running
	guard := make(chan struct{}, maxGoroutines)
	outChannel := make(chan *Imports, maxGoroutines)
	importSet := make(map[string]bool)
	var parsedImports Imports

	count := 0
	fileIndex := 0

	nFiles := len(filenames)
	parsedImports.FileCount = nFiles

	for count < nFiles {
		select {
		case guard <- struct{}{}:
			// Start a new goroutine and update counter
			if fileIndex >= nFiles {
				continue
			} else {
				go fromFileAsync(filenames[fileIndex], outChannel, guard)
				fileIndex += 1
			}

		case fileImports := <-outChannel:
			// Update import set
			for _, fileImport := range fileImports.ImportArray {
				moduleName := nameRegex.FindString(fileImport)
				// Get the module name from the import (retain scope, exclude module subpaths)
				importSet[moduleName] = true
			}
			// Update log arrays
			parsedImports.Logs.Debug = append(parsedImports.Logs.Debug, fileImports.Logs.Debug...)
			parsedImports.Logs.Warning = append(parsedImports.Logs.Warning, fileImports.Logs.Warning...)
			parsedImports.Logs.Err = append(parsedImports.Logs.Err, fileImports.Logs.Err...)
			parsedImports.Logs.Verbose = append(parsedImports.Logs.Verbose, fileImports.Logs.Verbose...)
			parsedImports.Logs.Info = append(parsedImports.Logs.Info, fileImports.Logs.Info...)
			count += 1
		}
	}

	// Transform the set into an array
	for k := range importSet {
		parsedImports.ImportArray = append(parsedImports.ImportArray, k)
	}

	return &parsedImports

}

func fromFileAsync(filename string, outChannel chan *Imports, finished chan struct{}) {
	outChannel <- FromFile(filename)
	<-finished
}

func FromFile(filename string) *Imports {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Imports{Logs: LogMap{Err: []string{fmt.Sprintf("Error reading: %s. Error message is: %s", filename, err.Error())}}}
	}

	var logMap LogMap
	var importMap []string

	log := NewLogMap(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)
	ext := filepath.Ext(filename)

	sourceFile := logger.Source{
		Index:          0,
		KeyPath:        logger.Path{Text: filename},
		PrettyPath:     filename,
		Contents:       string(data),
		IdentifierName: filename,
	}

	if ext == ".css" {
		importMap = FromCSS(&log, &sourceFile)
	} else {
		importMap = FromECMA(&log, &sourceFile, ext)

	}
	return &Imports{importMap, logMap, 1}
}

func FromCSS(log *logger.Log, sourceFile *logger.Source) []string {
	var importMap []string

	ast := css_parser.Parse(*log, *sourceFile, css_parser.Options{})

	for _, record := range ast.ImportRecords {
		if !isDependencyLocal(record.Path.Text) {
			importMap = append(importMap, record.Path.Text)
		}
	}
	return importMap
}

func FromECMA(log *logger.Log, sourceFile *logger.Source, ext string) []string {
	var importMap []string

	options := config.Options{Mode: config.ModeBundle}
	if ext == ".ts" || ext == ".tsx" {
		options.TS = config.TSOptions{
			Parse: true,
		}
	}

	if ext == ".jsx" || ext == ".tsx" {
		options.JSX = config.JSXOptions{
			Parse: true,
		}
	}

	ast, _ := js_parser.Parse(*log, *sourceFile, js_parser.OptionsFromConfig(&options))

	for _, record := range ast.ImportRecords {
		if !isDependencyLocal(record.Path.Text) {
			importMap = append(importMap, record.Path.Text)
		}
	}

	return importMap
}

func isDependencyLocal(dependency string) bool {
	return (strings.HasPrefix(dependency, ".") || strings.HasPrefix(dependency, "/") || strings.HasPrefix(dependency, "data:"))
}
