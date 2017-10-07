package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var inFiles = flag.String("input", "", "the glob pattern of input text files. Empty string uses stdin and stdout.")
var patternFlag = flag.String("pat", "", "the regular expresion searched within each line")
var pattern *regexp.Regexp
var replace = flag.String("repl", "", "replace text for pattern matched")

var debugFlag = flag.Bool("debug", false, "prints debug info to standard error")
var debugFile = ioutil.Discard

func main() {
	progName := "ledit"
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s is a line editor to tranform lines in text files.\n", progName)
		fmt.Fprintf(os.Stderr, "Matches of the pattern pat are replaced with the replacement string repl.\n")
		fmt.Fprintf(os.Stderr, "The syntax of the regular is here: https://github.com/google/re2/wiki/Syntax\n")
		fmt.Fprintf(os.Stderr, "Inside repl, $ signs are interpreted as in the link below.\n")
		fmt.Fprintf(os.Stderr, "https://golang.org/pkg/regexp/#Regexp.Expand\n")
		fmt.Fprintf(os.Stderr, "For instance $1 represents the text of the first submatch.\n")
		fmt.Fprintf(os.Stderr, "The output file has an \".%s\" appended to the end.\n\n", progName)
		flag.PrintDefaults()
	}

	flag.Parse()
	if *debugFlag {
		debugFile = os.Stderr
		fmt.Fprintln(debugFile, "Debug for ", progName)
	}
	pattern = regexp.MustCompile(*patternFlag)
	fmt.Fprintf(debugFile, "glob pattern of input files: \"%s\"\n", *inFiles)
	fmt.Fprintf(debugFile, "regular expression pattern : \"%s\"\n", pattern)
	fmt.Fprintf(debugFile, "replacement expression     : \"%s\"\n", *replace)
	editFiles()
}

func editFiles() {
	inFiles, err := filepath.Glob(*inFiles)
	if err != nil {
		panic("not valid glob pattern of input text files")
	}

	if len(inFiles) == 0 {
		inFiles = append(inFiles, "")
	}

	for _, filename := range inFiles {
		if strings.Compare(filename, "") == 0 {
			fmt.Fprintln(debugFile, "editing file: <stdin>")
		} else {
			fmt.Fprintln(debugFile, "editing file: ", filename)
		}
		editFile(filename)
	}
}

func setFiles(filename string) (*os.File, *os.File) {
	var inFile *os.File
	var outFile *os.File
	if strings.Compare(filename, "") == 0 {
		inFile = os.Stdin
		outFile = os.Stdout
		return inFile, outFile
	}

	inFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	outFile, err = os.Create(filename + ".ledit")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return inFile, outFile
}

func editFile(filename string) {
	inFile, outFile := setFiles(filename)
	defer inFile.Close()
	defer outFile.Close()

	var err error
	lineNum := 1
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		outLine := line
		if pattern.MatchString(line) {
			fmt.Fprintln(debugFile, "line: ", lineNum)
			fmt.Fprintln(debugFile, "matched      : ", line)

			outlLine := pattern.ReplaceAllString(line, *replace)
			fmt.Fprintln(debugFile, "replaced with: ", outlLine)
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
