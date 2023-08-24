package main

type Instruction interface {
	toString() string
	compile() []int32
	toBinaryExpression() []string
}
