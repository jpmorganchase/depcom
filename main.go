package main

import (
	"fmt"

	"time"

	"depcom/parse"

	"github.com/mattn/go-zglob"
)

func globMatches(dirPath string) {
	// TODO: exclude node_modules if possible (maybe manually?)
	matches, err := zglob.Glob(dirPath + "/**/*.{tsx,jsx,mjs,cjs,ts,js,css}")
	var importSet []string

	start := time.Now()
	if err != nil {
		fmt.Println(err)
	} else {
		importSet = parse.FromFiles(matches)
	}
	elapsed := time.Since(start)
	fmt.Printf("Parsing took %s\n", elapsed)
	fmt.Printf("Global imports: %v\n", importSet)
}

func main() {
	//globMatches("./examples")
	globMatches("../modular/packages/modular-scripts/src")
}
