package main

import (
	"bufio"
	"os"
	"strings"
)

var declarations = map[string]bool{}

func ScanDeclarations(source *os.File) {
	reader := bufio.NewReader(source)
	line, readErr := reader.ReadString('\n')
	for ; readErr == nil; line, readErr = reader.ReadString('\n') {
		if strings.HasPrefix(line, "label") {
			parts := strings.Split(line, " ")
			declarations[parts[1]] = true
		}
	}
	if len(line) > 0 {
		if strings.HasPrefix(line, "label") {
			parts := strings.Split(line, " ")
			declarations[parts[1]] = true
		}
	}
}
