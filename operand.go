package main

import "fmt"

// Operand represents the type of operands used in the RPN calculator.
type Operand int8

// this represent the priority of the operands
const (
	LEFT_PARENTHESIS  Operand = -1
	RIGHT_PARENTHESIS Operand = -2
	ADD               Operand = 0
	SUB               Operand = 1
	MUL               Operand = 2
	DIV               Operand = 3
)

// ToOperand converts a string representation of an operand to its Operand type.
// It returns an error if the string does not match any known operand.
// The operands are defined as follows:
// ADD: "+", SUB: "-", MUL: "*", DIV: "/", LEFT_PARENTHESIS: "(", RIGHT_PARENTHESIS: ")"
//
// Example usage:
//
//	op, err := ToOperand("+")
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//	    fmt.Println(op) // Output: ADD
//	}
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

// IsPriorityThan checks if the current operand has a higher priority than the other operand.
// The priority is defined as follows:
// - MUL and DIV have higher priority than ADD and SUB.
// - If both operands are of the same type (MUL or DIV, or ADD or SUB), they are considered equal in priority.
// - The evaluation is left-to-right for operands of the same priority.
//
// Example usage:
//
//	op1 := ADD
//	op2 := MUL
//	fmt.Println(op1.IsPriorityThan(op2)) // Output: false
//	fmt.Println(op2.IsPriorityThan(op1)) // Output: true
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

// String returns the string representation of the Operand.
// It returns "+" for ADD, "-" for SUB, "*" for MUL, "/" for DIV, and "unknown" for any other value.
//
// Example usage:
//
//	op := ADD
//	fmt.Println(op.String()) // Output: "+"
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

// CalculateString performs the calculation based on the operand and two float64 values.
// It returns the result of the operation or an error if the operation is invalid.
// The operands are defined as follows:
// ADD: addition, SUB: subtraction, MUL: multiplication, DIV: division
//
// Example usage:
//
//	result, err := ADD.CalculateString(3.0, 5.0)
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//	    fmt.Println(result) // Output: 8.0
//	}
//
// If division by zero is attempted, it returns an error.
// If an unknown operand is used, it returns an error.
// If the operand is not recognized, it returns an error.
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
