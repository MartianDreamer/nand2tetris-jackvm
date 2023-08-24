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
	Translate(os.Args[1])
}
