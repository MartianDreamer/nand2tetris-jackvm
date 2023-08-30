package main

import (
	"bufio"
	"os"
	"strings"
)

func Translate(source *os.File, filePath string) {
	reader := bufio.NewReader(source)
	outputPath := strings.Replace(filePath, ".vm", ".asm", 1)
	filePathParts := strings.Split(filePath, "/")
	scope := strings.Replace(filePathParts[len(filePathParts)-1], ".vm", "", 1)
	outputFile := createFile(outputPath)
	writer := bufio.NewWriter(outputFile)
	defer func(outputFile *os.File) {
		errFlush := writer.Flush()
		if errFlush != nil {
			panic("failed to write file")
		}
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)
	write(writer, setupSp())
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
	write(writer, endProgram())
}

func endProgram() string {
	return "(END_PROG)\n" +
		"@END_PROG\n" +
		"0; JMP\n"
}

func setupSp() string {
	return "@256\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n"
}
