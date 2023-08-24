package main

type ArithmeticInstruction struct {
	instruction string
}

func (receiver ArithmeticInstruction) toString() string {
	return receiver.instruction
}

func (receiver ArithmeticInstruction) compile() []int32 {
	return []int32{0, 0, 0}
}

func (receiver ArithmeticInstruction) toBinaryExpression() []string {
	return []string{"1111", "0111", "1011"}
}
