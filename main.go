package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"depcom/parse"

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

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	} else if directory != "" {
		parsedImports = globMatches(directory)

	} else {
		tail := flag.Args()
		parsedImports = parse.FromFiles(tail)
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
	} else {
		return parse.FromFiles(matches)
	}
}
