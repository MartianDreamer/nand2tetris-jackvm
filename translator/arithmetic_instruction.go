package translator

func getTranslationFunction(instruction string) func() []string {
	switch instruction {
	case "add":
		return add
	default:
		panic("invalid instruction")
	}
}

type arithmeticInstruction struct {
	instruction string
}

func (receiver arithmeticInstruction) toString() string {
	return receiver.instruction
}

func (receiver arithmeticInstruction) Compile() []string {
	result := startInstruction()
	result = append(result, getTranslationFunction(receiver.instruction)()...)
	result = append(result, endInstruction()...)
	return result
}

func isValid(s string) bool {
	switch s {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return true
	default:
		return false
	}
}

func startInstruction() []string {
	return []string{
		"@SP",
		"A=M-1",
		"D=M",
	}
}

func endInstruction() []string {
	return []string{
		"D=A+1",
		"@SP",
		"M=D",
	}
}

func add() []string {
	return []string{
		"A=A-1",
		"M=M+D",
	}
}
