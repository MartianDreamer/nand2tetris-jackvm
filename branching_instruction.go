package main

import "strings"

func isValidBranchingOperation(parts []string) bool {
	return (len(parts) == 1 && (strings.HasPrefix(parts[0], "goto") || strings.HasPrefix(parts[0], "if-goto"))) ||
		(len(parts) == 2 && strings.HasPrefix(parts[0], "label"))
}

type branchingInstruction struct {
	instruction string
	label       *string
}

func (b branchingInstruction) Compile() string {
	//TODO implement me
	panic("implement me")
}
