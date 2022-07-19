package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cristiano-belloni/depcom/parse"

	"github.com/mattn/go-zglob"
)

type DepcomResult struct {
	Time string
	parse.Imports
}

func main() {
	var directory string
	var parsedImports *parse.Imports
	var help bool

	start := time.Now()

	flag.StringVar(&directory, "d", "", "Directory to glob")
	flag.BoolVar(&help, "h", false, "Display help")
	flag.Parse()
	tail := flag.Args()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if directory != "" {
		parsedImports = globMatches(directory)

	} else if len(tail) > 0 {
		parsedImports = parse.FromFiles(tail)
	} else {
		fmt.Println("No arguments specified, showing help.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	elapsed := time.Since(start)

	jsonResult, _ := json.Marshal(DepcomResult{Time: elapsed.String(), Imports: *parsedImports})
	fmt.Println(string(jsonResult))
}

func globMatches(dirPath string) *parse.Imports {
	globExpression := dirPath + "/**/*.{tsx,jsx,mjs,cjs,ts,js,css}"
	matches, err := zglob.Glob(globExpression)

	if err != nil {
		return &parse.Imports{Logs: parse.LogMap{Err: []string{fmt.Sprintf("Error globbing: %s. Error message is: %s", globExpression, err.Error())}}}
	} else if len(matches) == 0 {
		return &parse.Imports{Logs: parse.LogMap{Err: []string{fmt.Sprintf("No matches found globbing: %s", globExpression)}}}

	} else {
		return parse.FromFiles(matches)
	}
}
