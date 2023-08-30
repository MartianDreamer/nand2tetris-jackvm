package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("compiled file not found")
	} else if !strings.HasSuffix(args[1], ".vm") {
		panic("invalid file type")
	}
	source := openFile(os.Args[1])
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			return
		}
	}(source)
	ScanDeclarations(source)
	Translate(source, os.Args[1])

}
