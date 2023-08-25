package main

import "strconv"

var count uint64 = 0

func getTranslationFunction(instruction string) func() string {
	switch instruction {
	case "add":
		return add
	case "sub":
		return sub
	case "neg":
		return neg
	case "and":
		return and
	case "or":
		return or
	case "not":
		return not
	case "gt":
		return gt
	case "lt":
		return lt
	case "eq":
		return eq
	default:
		panic("invalid instruction " + instruction)
	}
}

type arithmeticInstruction struct {
	instruction string
}

func (receiver arithmeticInstruction) Compile() string {
	return startInstruction() + getTranslationFunction(receiver.instruction)() + endInstruction()
}

func isValidArithmeticOperations(s string) bool {
	switch s {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return true
	default:
		return false
	}
}

func startInstruction() string {
	return "@SP\n" +
		"A=M-1\n"
}

func endInstruction() string {
	return "D=A+1\n" +
		"@SP\n" +
		"M=D\n"
}

func add() string {
	return "D=M\n" +
		"A=A-1\n" +
		"M=M+D\n"
}

func sub() string {
	return "D=M\n" +
		"A=A-1\n" +
		"M=M-D\n"
}

func neg() string {
	return "M=-M\n"
}

func and() string {
	return "D=M\n" +
		"A=A-1\n" +
		"M=D&M\n"
}

func or() string {
	return "D=M\n" +
		"A=A-1\n" +
		"M=D|M\n"
}

func not() string {
	return "M=!M\n"
}

func gt() string {
	result := "D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@NOT_TRUE" + strconv.FormatUint(count, 10) + "\n" +
		"D; JLE\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=-1\n" +
		"@END" + strconv.FormatUint(count, 10) + "\n" +
		"0;JMP\n" +
		"(NOT_TRUE" + strconv.FormatUint(count, 10) + ")\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=0\n" +
		"(END" + strconv.FormatUint(count, 10) + ")\n" +
		"A=D\n"
	count++
	return result
}

func lt() string {
	result := "D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@NOT_TRUE" + strconv.FormatUint(count, 10) + "\n" +
		"D; JGE\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=-1\n" +
		"@END" + strconv.FormatUint(count, 10) + "\n" +
		"0;JMP\n" +
		"(NOT_TRUE" + strconv.FormatUint(count, 10) + ")\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=0\n" +
		"(END" + strconv.FormatUint(count, 10) + ")\n" +
		"A=D\n"
	count++
	return result
}

func eq() string {
	result := "D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@NOT_TRUE" + strconv.FormatUint(count, 10) + "\n" +
		"D; JNE\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=-1\n" +
		"@END" + strconv.FormatUint(count, 10) + "\n" +
		"0;JMP\n" +
		"(NOT_TRUE" + strconv.FormatUint(count, 10) + ")\n" +
		"@SP\n" +
		"D=M\n" +
		"@2\n" +
		"AD=D-A\n" +
		"M=0\n" +
		"(END" + strconv.FormatUint(count, 10) + ")\n" +
		"A=D\n"
	count++
	return result
}
