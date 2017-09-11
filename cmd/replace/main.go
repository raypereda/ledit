package main

import (
	"fmt"
	"flag"
	"path/filepath"
)


var pattern = flag.String("pattern", "", "line pattern for searching for")
var rline = flag.String("rline", "", "replace line for matched lines")
var inGlob = flag.String("inglob", "", "the glob pattern of input text files")

func main() {
	flag.Parse()
	fmt.Println("pattern:", *pattern)
	fmt.Println("rline: ", *rline)
	infiles, err := filepath.Glob(*inGlob)
	if err != nil {
		panic("not value Glob pattern")
	}

	fmt.Println("infiles: ", infiles)

	fmt.Println("hello")
}