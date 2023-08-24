package main

import "strconv"

type memoryInstructionType int8
type memorySegment int8

const (
	push memoryInstructionType = iota
	pop
)

func getMemoryInstructionType(s string) memoryInstructionType {
	switch s {
	case "push":
		return push
	case "pop":
		return pop
	default:
		return invalid
	}
}

const (
	arg memorySegment = iota
	lcl
	this
	that
	constant
	static
	pointer
	temp
)

func getSegment(s string) memorySegment {
	switch s {
	case "argument":
		return arg
	case "local":
		return lcl
	case "this":
		return this
	case "that":
		return that
	case "constant":
		return constant
	case "static":
		return static
	case "pointer":
		return pointer
	case "temp":
		return temp
	default:
		return invalid
	}
}

type memoryInstruction struct {
	scope   string
	iType   memoryInstructionType
	segment memorySegment
	offset  uint8
}

func (m memoryInstruction) Compile() string {
	return "hello"
}

func pushToStack() string {
	return "@SP\n" +
		"A=M\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M+1\n"
}

func popFromStack() string {
	return "@SP\n" +
		"A=M\n" +
		"D=M\n" +
		"@R13\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M-1\n"
}

func assignD(offset int16) string {
	return "@" + strconv.FormatInt(int64(offset), 0) + "\n" +
		"D=A\n"
}

func getValueFromArg(offset uint8) string {
	return assignD(int16(offset)) +
		"@ARG\n" +
		"A=M\n" +
		"A=D+A\n" +
		"D=A\n"
}
