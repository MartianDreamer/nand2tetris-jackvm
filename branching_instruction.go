package main

import "strings"

func isValidBranchingOperation(parts []string) bool {
	return len(parts) == 2 &&
		(strings.Compare(parts[0], "label") == 0 ||
			((strings.Compare(parts[0], "goto") == 0 ||
				strings.Compare(parts[0], "if-goto") == 0) &&
				declarations[parts[1]]))
}

type branchingInstruction struct {
	instruction string
	label       string
}

func (b branchingInstruction) Compile() string {
	switch b.instruction {
	case "label":
		return "(" + b.label + ")\n"
	case "goto":
		return "@" + b.label + "\n" +
			"0;JMP\n"
	case "if-goto":
		return "@SP\n" +
			"A=M\n" +
			"D=M+1\n" +
			"@" + b.label + "\n" +
			"D;JEQ\n"
	default:
		panic("failed to translate this instruction " + b.instruction + " " + b.label)
	}
}
