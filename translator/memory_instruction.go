package translator

import "errors"

type memoryInstructionType int8
type memorySegment int8

const (
	PUSH memoryInstructionType = iota
	POP
)

func getMemoryInstructionType(s string) memoryInstructionType {
	switch s {
	case "push":
		return PUSH
	case "pop":
		return POP
	default:
		panic("invalid instruction")
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
	invalid
)

func getSegment(s string) (memorySegment, error) {
	switch s {
	case "argument":
		return arg, nil
	case "local":
		return lcl, nil
	case "this":
		return this, nil
	case "that":
		return that, nil
	case "constant":
		return constant, nil
	case "static":
		return static, nil
	case "pointer":
		return pointer, nil
	case "temp":
		return temp, nil
	default:
		return invalid, errors.New("invalid segment")
	}
}

type memoryInstruction struct {
	instruction string
	iType       memoryInstructionType
	segment     memorySegment
	offset      uint8
}

func (m memoryInstruction) toString() string {
	return m.instruction
}

func (m memoryInstruction) Compile() []string {
	return []string{}
}
