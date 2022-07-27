package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/jpmorganchase/depcom/parse"

	"github.com/mattn/go-zglob"
)

type DepcomResult struct {
	Time string
	parse.Imports
}

type ArrayFlags []string

func (i *ArrayFlags) String() string {
	return ""
}

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var directoryPath string
	var includePattern string
	var excludePattern ArrayFlags
	var parsedImports *parse.Imports
	var help bool

	flag.StringVar(&directoryPath, "d", "./", "Base directory path")
	flag.StringVar(&includePattern, "a", "/**/*.{tsx,jsx,mjs,cjs,ts,js,css}", "Glob pattern of files to analyze")
	flag.Var(&excludePattern, "x", "Glob pattern of files to exclude from analysis")
	flag.BoolVar(&help, "h", false, "Display help")
	flag.Parse()
	tail := flag.Args()

	cleanBaseDirectory := path.Clean(directoryPath)

	start := time.Now()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	} else if !isFlagPassed("d") && !isFlagPassed("a") && !isFlagPassed("x") && len(tail) > 0 {
		parsedImports = parse.FromFiles(tail)
	} else {
		parsedImports = globMatches(cleanBaseDirectory, includePattern, excludePattern)

	}

	elapsed := time.Since(start)

	jsonResult, _ := json.Marshal(DepcomResult{Time: elapsed.String(), Imports: *parsedImports})
	fmt.Println(string(jsonResult))
}

func globMatches(baseDirectory string, includePattern string, excludePatterns []string) *parse.Imports {
	exclusions := [][]string{}
	includedFiles, err := zglob.Glob(baseDirectory + "/" + includePattern)

	if err != nil {
		return &parse.Imports{Logs: parse.LogMap{Err: []string{fmt.Sprintf("Error globbing files: %s. Error message is: %s", includePattern, err.Error())}}}
	}

	for _, excludePattern := range excludePatterns {
		if excludePattern != "" {
			exclusionMatches, err := zglob.Glob(baseDirectory + "/" + excludePattern)
			if err != nil {
				return &parse.Imports{Logs: parse.LogMap{Err: []string{fmt.Sprintf("Error globbing excluded files: %s. Error message is: %s", excludePattern, err.Error())}}}
			}
			exclusions = append(exclusions, exclusionMatches)
		}
	}

	fileList := multipleArrayDifference(includedFiles, exclusions)

	if len(fileList) == 0 {
		return &parse.Imports{Logs: parse.LogMap{Err: []string{"No matches found"}}}
	} else {
		return parse.FromFiles(fileList)
	}
}

func multipleArrayDifference(a []string, b [][]string) []string {
	mb := make(map[string]struct{}, len(b))

	// Accumulate all the exclusions in a map
	for _, x := range b {
		for _, y := range x {
			mb[y] = struct{}{}
		}
	}
	var diff []string

	// Everything that's not in the accumulator gets added
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
