package main

func getTranslationFunction(instruction string) func() []string {
	switch instruction {
	case "add":
		return add
	case "sub":
		return sub
	case "neg":
		return neg
	case "and":
		return and
	case "or":
		return or
	case "not":
		return not
	case "qt":
		return gt
	case "lt":
		return lt
	case "eq":
		return eq
	default:
		panic("invalid instruction")
	}
}

type arithmeticInstruction struct {
	instruction string
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
		"D=M",
		"A=A-1",
		"M=M+D",
	}
}

func sub() []string {
	return []string{
		"D=M",
		"A=A-1",
		"M=M-D",
	}
}

func neg() []string {
	return []string{
		"M=-M",
	}
}

func and() []string {
	return []string{
		"D=M",
		"A=A-1",
		"M=D&M",
	}
}

func or() []string {
	return []string{
		"D=M",
		"A=A-1",
		"M=D|M",
	}
}

func not() []string {
	return []string{
		"M=!M",
	}
}

func gt() []string {
	return []string{
		"D=M",
		"A=A-1",
		"M=D-M",
	}
}

func lt() []string {
	return []string{
		"D=M",
		"A=A-1",
		"M=M-D",
	}
}

func eq() []string {
	return []string{
		"D=M",
		"A=A-1",
		"D=M-D",
		"@SP",
		"A=M",
		"M=!D",
		"A=A-1",
		"D=M",
		"A=A-1",
		"M=D-M",
		"M=!M",
		"@SP",
		"A=M",
		"D=M",
		"A=A-1",
		"A=A-1",
		"M=D&M",
	}
}
