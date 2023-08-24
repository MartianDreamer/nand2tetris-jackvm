package main

import (
	"bufio"
	"os"
	"strings"
)

func Translate(filename string) {
	source := openFile(filename)
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			return
		}
	}(source)
	reader := bufio.NewReader(source)
	outputName := strings.Replace(filename, ".vm", ".asm", 1)
	scope := strings.Replace(filename, ".vm", "", 1)
	outputFile := createFile(outputName)
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)
	writer := bufio.NewWriter(outputFile)
	line, readErr := reader.ReadString('\n')
	for readErr != nil {
		instruction := Parse(line, scope)
		write(writer, instruction.Compile())
		line, readErr = reader.ReadString('\n')
	}
	if len(line) > 0 {
		instruction := Parse(line, scope)
		write(writer, instruction.Compile())
	}
}
