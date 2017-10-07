package main

import (
	"flag"
	"os"
	"testing"
)

func TestSetFiles(t *testing.T) {
	inFile, outFile := setFiles("")
	if inFile != os.Stdin {
		t.Error(`setFiles("") does not set inFile to Stdin`)
	}
	if outFile != os.Stdout {
		t.Error(`setFiles("") does not set outFile to Stdout`)
	}
}

func ExampleMain() {
	flag.Set("input", "test_input1.txt")
	flag.Set("pat", "aaa")
	flag.Set("repl", "bbbbbb")
	flag.Set("debug", "1")

	*debugFlag = false
    debugFile = os.Stdout
	// testableMain()
	main()

	// Output:
	// glob pattern of input files: "test_input1.txt"
	// regular expression pattern : "aaa"
	// replacement expression     : "bbbbbb"
	// editing file:  test_input1.txt
	// line:  2
	// matched      :  second line with aaa
	// replaced with:  second line with bbbbbb
}
