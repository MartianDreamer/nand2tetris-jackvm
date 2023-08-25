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
	offset  int
}

func (m memoryInstruction) Compile() string {
	switch m.iType {
	case push:
		return compilePushInstruction(m)
	case pop:
		return compilePopFunction(m)
	default:
		panic("invalid instruction")
	}
}

func assignD(offset int) string {
	return "@" + strconv.FormatInt(int64(offset), 10) + "\n" +
		"D=A\n"
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
		"AM=M-1\n" +
		"D=M\n" +
		"@R13\n" +
		"A=M\n" +
		"M=D\n"
}

func getValueFromArg(offset int) string {
	return assignD(offset) +
		"@ARG\n" +
		"A=D+M\n" +
		"D=M\n"
}

func getValueFromLcl(offset int) string {
	return assignD(int(offset)) +
		"@LCL\n" +
		"A=D+M\n" +
		"D=M\n"
}

func getValueFromThis(offset int) string {
	return assignD(offset) +
		"@THIS\n" +
		"A=D+M\n" +
		"D=M\n"
}

func getValueFromThat(offset int) string {
	return assignD(offset) +
		"@THAT\n" +
		"A=D+M\n" +
		"D=M\n"
}

func getValueFromConst(offset int) string {
	return assignD(offset)
}

func getValueFromStatic(offset int, scope string) string {
	return "@" + scope + "." + strconv.FormatInt(int64(offset), 10) + "\n" +
		"D=M\n"
}

func getValueFromTemp(offset int) string {
	return "@" + strconv.FormatInt(int64(5+offset), 10) + "\n" +
		"D=M\n"
}

func getValueFromPointer(offset int) string {
	switch offset {
	case 0:
		return "@THIS\n" +
			"D=M\n"
	case 1:
		return "@THAT\n" +
			"D=M\n"
	default:
		panic("invalid offset in pointer segment")
	}
}

func compilePushInstruction(m memoryInstruction) string {
	var result string
	switch m.segment {
	case arg:
		result = getValueFromArg(m.offset)
	case lcl:
		result = getValueFromLcl(m.offset)
	case this:
		result = getValueFromThis(m.offset)
	case that:
		result = getValueFromThat(m.offset)
	case static:
		result = getValueFromStatic(m.offset, m.scope)
	case constant:
		result = getValueFromConst(m.offset)
	case temp:
		result = getValueFromTemp(m.offset)
	case pointer:
		result = getValueFromPointer(m.offset)
	default:
		panic("invalid segment")
	}
	result += pushToStack()
	return result
}

func findArgAddress(offset int) string {
	return assignD(offset) +
		"@ARG\n" +
		"D=D+M\n" +
		"@R13\n" +
		"M=D\n"
}

func findLclAddress(offset int) string {
	return assignD(offset) +
		"@LCL\n" +
		"D=D+M\n" +
		"@R13\n" +
		"M=D\n"
}

func findThisAddress(offset int) string {
	return assignD(offset) +
		"@THIS\n" +
		"D=D+M\n" +
		"@R13\n" +
		"M=D\n"
}

func findThatAddress(offset int) string {
	return assignD(offset) +
		"@THAT\n" +
		"D=D+M\n" +
		"@R13\n" +
		"M=D\n"
}

func putOnPointerSegment(offset int) string {
	switch offset {
	case 0:
		return "@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@THIS\n" +
			"M=D\n"
	case 1:
		return "@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@THAT\n" +
			"M=D\n"
	default:
		panic("invalid offset in pointer segment")
	}
}

func putOnTempSegment(offset int) string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@" + strconv.FormatInt(int64(5+offset), 10) + "\n" +
		"M=D\n"
}

func putOnStaticSegment(offset int, scope string) string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@" + scope + "." + strconv.FormatInt(int64(offset), 10) + "\n" +
		"M=D\n"
}

func compilePopFunction(m memoryInstruction) string {
	switch m.segment {
	case arg:
		return findArgAddress(m.offset) + popFromStack()
	case lcl:
		return findLclAddress(m.offset) + popFromStack()
	case this:
		return findThisAddress(m.offset) + popFromStack()
	case that:
		return findThatAddress(m.offset) + popFromStack()
	case static:
		return putOnStaticSegment(m.offset, m.scope)
	case temp:
		return putOnTempSegment(m.offset)
	case pointer:
		return putOnPointerSegment(m.offset)
	default:
		panic("invalid segment")
	}
}
