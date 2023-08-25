package main

import (
	"bufio"
	"os"
	"strings"
)

func Translate(filePath string) {
	filePath = strings.Replace(filePath, "\\", "/", -1)
	source := openFile(filePath)
	reader := bufio.NewReader(source)
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			return
		}
	}(source)
	outputPath := strings.Replace(filePath, ".vm", ".asm", 1)
	filePathParts := strings.Split(filePath, "/")
	scope := strings.Replace(filePathParts[len(filePathParts)-1], ".vm", "", 1)
	outputFile := createFile(outputPath)
	writer := bufio.NewWriter(outputFile)
	defer func(outputFile *os.File) {
		writer.Flush()
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)
	line, readErr := reader.ReadString('\n')
	for ; readErr == nil; line, readErr = reader.ReadString('\n') {
		line = strings.Replace(line, "\r\n", "", 1)
		line = strings.Replace(line, "\n", "", 1)
		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}
		instruction := Parse(line, scope)
		write(writer, instruction.Compile())
	}
	if len(line) > 0 {
		instruction := Parse(line, scope)
		write(writer, instruction.Compile())
	}
}
