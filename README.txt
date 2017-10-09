
From the command-line help:

$ ledit -h
ledit is a line editor to transform lines in text files.
Matches of the pattern pat are replaced with the replacement string repl.
The syntax of the regular is here: https://github.com/google/re2/wiki/Syntax
Inside repl, $ signs are interpreted as in the link below.
https://golang.org/pkg/regexp/#Regexp.Expand
For instance $1 represents the text of the first submatch.
The output file has an ".ledit" appended to the end.

  -debug
    	prints debug info to standard error
  -input string
    	the glob pattern of input text files. Empty string uses stdin and stdout.
  -pat string
    	the regular expresion searched within each line
  -repl string
    	replace text for pattern matched
      
-------------------
Examples:

Example #1
$ ledit -debug -pat=dog -repl=cat -input=input1.txt
Debug for  ledit
glob pattern of input files: "input1.txt"
regular expression pattern : "dog"
replacement expression     : "cat"
editing file:  input1.txt
line:  2
matched      :  Let's go walk the dog at night.
replaced with:  Let's go walk the cat at night.

Example #2
$ ledit -debug -pat="<tag39>false" -repl="<tag39>true" -input=input2.html
Debug for  ledit
glob pattern of input files: "input2.html"
regular expression pattern : "<tag39>false"
replacement expression     : "<tag39>true"
editing file:  input2.html
line:  4
matched      :  <tag39>false</tag39>
replaced with:  <tag39>true</tag39>

Example #3
$ ledit -debug -pat=10.0.0.2 -repl=10.0.0.99 -input=input3.toml
Debug for  ledit
glob pattern of input files: "input3.toml"
regular expression pattern : "10.0.0.2"
replacement expression     : "10.0.0.99"
editing file:  input3.toml
line:  4
matched      :    ip = "10.0.0.2"
replaced with:    ip = "10.0.0.99"

The OSX executable is cmd/ledit/ledit.
The Windows 64-bit executable is cmd/ledit/ledit.exe.

sed is a stream editor that is a bit old and cryptic. 
You can problem solve your problem with a line editing using sed. 
I prefer something easy to use and easy to modify source code.
