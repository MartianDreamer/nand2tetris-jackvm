package main

import (
	. "bufio"
	"os"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open " + filename)
	}
	return file
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic("failed to create file " + filename)
	}
	return file
}

func write(writer *Writer, content string) {
	_, err := writer.WriteString(content)
	if err != nil {
		panic("failed to write file")
	}
}
