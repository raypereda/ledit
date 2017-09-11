package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var patternFlag = flag.String("pattern", "", "line pattern for searching for")
var rline = flag.String("rline", "", "replace line for matched lines")
var inGlob = flag.String("inglob", "", "the glob pattern of input text files")
var pattern = regexp.MustCompile(*patternFlag)

func main() {
	flag.Parse()
	fmt.Println("pattern:", *patternFlag)
	fmt.Println("rline: ", *rline)
	fmt.Println("pattern: ", pattern)

	editFiles()
	fmt.Println("hello")
}

func editFiles() {
	infiles, err := filepath.Glob(*inGlob)
	if err != nil {
		panic("not value Glob pattern")
	}
	fmt.Println("infiles: ", infiles)

	for _, filename := range infiles {
		fmt.Println("editing file: ", filename)
		editFile(filename)
	}
}

func editFile(filename string) {
	// open a file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// make sure it gets closed
	defer file.Close()

	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			fmt.Print("MATCH: ")
		}
		fmt.Println(line)
	}

	// check for errors
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
