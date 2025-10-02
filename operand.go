package main

import "fmt"

type Operand int8

const (
	LEFT_PARENTHESIS  Operand = -1
	RIGHT_PARENTHESIS Operand = -2
	ADD               Operand = 0
	SUB               Operand = 1
	MUL               Operand = 2
	DIV               Operand = 3
)

func ToOperand(exp any) (Operand, error) {
	switch exp {
	case "+":
		return ADD, nil
	case "-":
		return SUB, nil
	case "*":
		return MUL, nil
	case "/":
		return DIV, nil
	case "(":
		return LEFT_PARENTHESIS, nil
	case ")":
		return RIGHT_PARENTHESIS, nil
	default:
		return -1, fmt.Errorf("error: this string %v is not operand", exp)
	}
}

func (o Operand) IsPriorityThan(other Operand) bool {
	// MUL and DIV have higher priority than ADD and SUB
	if (o == MUL || o == DIV) && (other == ADD || other == SUB) {
		return true
	}

	// Same priority level, left-to-right evaluation
	if (o == MUL || o == DIV) && (other == MUL || other == DIV) {
		return true
	}
	if (o == ADD || o == SUB) && (other == ADD || other == SUB) {
		return true
	}

	return false
}

func (o Operand) String() string {
	switch o {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	default:
		return "unknown"
	}
}

func (o Operand) CalculateString(a, b float64) (float64, error) {
	switch o {
	case ADD:
		return a + b, nil
	case SUB:
		return a - b, nil
	case MUL:
		return a * b, nil
	case DIV:
		if b == 0 {
			return 0, fmt.Errorf("error: division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("error: unknown operand %s", o)
	}
}
