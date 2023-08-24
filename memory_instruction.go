package main

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

func (m memoryInstruction) Compile() []string {
	return []string{}
}
