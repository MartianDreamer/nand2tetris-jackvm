package translator

import (
	"strconv"
	"strings"
)

const (
	separator        = " "
	arithmetic uint8 = 1
	memory     uint8 = 2
)

type Instruction interface {
	toString() string
	Compile() []string
}

func Parse(instruction string) Instruction {
	parts := strings.Split(instruction, separator)
	category := categorizeInstruction(parts)
	switch category {
	case arithmetic:
		return arithmeticInstruction{
			instruction: instruction,
		}
	case memory:
		offset, _ := strconv.ParseUint(parts[2], 0, 8)
		segment, _ := getSegment(parts[1])
		return memoryInstruction{
			instruction: instruction,
			iType:       getMemoryInstructionType(parts[0]),
			segment:     segment,
			offset:      uint8(offset),
		}
	default:
		panic("failed to parse instruction")
	}
}

func categorizeInstruction(parts []string) uint8 {
	if len(parts) == 3 &&
		(strings.Compare(parts[0], "push") == 0 || strings.Compare(parts[0], "pop") == 0) {
		_, err := getSegment(parts[1])
		if err != nil {
			panic("invalid segment")
		}
		return arithmetic
	} else if len(parts) == 1 &&
		isValid(parts[0]) {
		return memory
	}
	return 0
}
