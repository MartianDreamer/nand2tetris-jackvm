package main

import (
	"strconv"
	"strings"
)

type validatorFunc func(string) bool

var validators = map[string]validatorFunc{
	"temp": func(s string) bool {
		parts := strings.Split(s, " ")
		address, _ := strconv.Atoi(parts[2])
		return address >= 5 && address <= 12
	},
}

func Validate(statement string) bool {
	statementType := getStatementType(statement)
	validator := validators[statementType]
	if validator != nil {
		return validator(statement)
	}
	return true
}

func getStatementType(statement string) string {
	parts := strings.Split(statement, " ")
	return parts[0]
}
