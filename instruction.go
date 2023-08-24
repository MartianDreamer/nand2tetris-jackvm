package main

import (
	"strconv"
	"strings"
)

const (
	separator       = " "
	invalid         = -1
	arithmetic int8 = 1
	memory     int8 = 2
)

type Instruction interface {
	Compile() string
}

func Parse(instruction string, scope string) Instruction {
	parts := strings.Split(instruction, separator)
	category := categorizeInstruction(parts)
	switch category {
	case arithmetic:
		return arithmeticInstruction{
			instruction: instruction,
		}
	case memory:
		offset, _ := strconv.ParseUint(parts[2], 0, 8)
		segment := getSegment(parts[1])
		return memoryInstruction{
			scope:   scope,
			iType:   getMemoryInstructionType(parts[0]),
			segment: segment,
			offset:  uint8(offset),
		}
	default:
		panic("failed to parse instruction")
	}
}

func categorizeInstruction(parts []string) int8 {
	if len(parts) == 3 &&
		getMemoryInstructionType(parts[0]) != invalid &&
		getSegment(parts[1]) != invalid {
		return arithmetic
	} else if len(parts) == 1 &&
		isValid(parts[0]) {
		return memory
	}
	return invalid
}
