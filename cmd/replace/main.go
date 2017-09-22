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
var pattern = regexp.MustCompile(*patternFlag)
var replace = flag.String("replace", "", "replace line for matched lines")
var inFiles = flag.String("infiles", "", "the glob pattern of input text files")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Replace is a tool for replacing lines in text files.\n")
		fmt.Fprintf(os.Stderr, "Lines matching a regular expression are replaced with fixed text.\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()
	pattern = regexp.MustCompile(*patternFlag)	
	fmt.Println("pattern: \"", pattern, "\"")	
	fmt.Println("replace: \"", *replace, "\"")
	fmt.Println("infiles: \"", *inFiles, "\"")
	editFiles()
}

func editFiles() {
	inFiles, err := filepath.Glob(*inFiles)
	if err != nil {
		panic("not valid glob pattern of input text files")
	}
	fmt.Println("infiles: ", inFiles)

	for _, filename := range inFiles {
		fmt.Println("editing file: ", filename)
		editFile(filename)
	}
}

func editFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()

	outFile, err := os.Create(filename + ".edited")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer outFile.Close()

	lineNum := 1
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		outLine := line
		if pattern.MatchString(line) {
			fmt.Printf("line: %3d matched: %s\n", lineNum, line) 
			
			outLine = *replace
			fmt.Println("replaced: ", outLine)	
		}
		_, err := fmt.Fprintln(outFile, outLine)
		if err != nil {
			panic(err)
		}
		lineNum++
	}

	// check for errors
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
