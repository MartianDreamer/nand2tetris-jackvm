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
	branching  int8 = 3
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
		offset, _ := strconv.ParseInt(parts[2], 0, 16)
		segment := getSegment(parts[1])
		return memoryInstruction{
			scope:   scope,
			iType:   getMemoryInstructionType(parts[0]),
			segment: segment,
			offset:  int(offset),
		}
	case branching:
		return branchingInstruction{
			instruction: parts[0],
			label:       parts[1],
		}
	default:
		panic("failed to parse instruction " + instruction + " in " + scope)
	}
}

func categorizeInstruction(parts []string) int8 {
	if len(parts) == 3 &&
		getMemoryInstructionType(parts[0]) != invalid &&
		getSegment(parts[1]) != invalid {
		return memory
	} else if len(parts) == 1 &&
		isValidArithmeticOperations(parts[0]) {
		return arithmetic
	} else if isValidBranchingOperation(parts) {
		return branching
	}
	return invalid
}
