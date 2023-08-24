package main

type MemoryInstructionType int8
type MemorySegment int8

const (
	PUSH MemoryInstructionType = iota
	POP
)

const (
	ARG MemorySegment = iota
	LCL
	THIS
	THAT
	CONST
	STATIC
	POINTER
	TEMP
)

type MemoryInstruction struct {
	instruction string
	iType       MemoryInstructionType
	segment     MemorySegment
	offset      int8
}
