package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RPN represents a Reverse Polish Notation calculator
// It contains two stacks: one for numbers and one for operators.
// It also contains the expression in an array format and its size.
// The expression is expected to be a valid mathematical expression with full brackets.
// The expression should be in the form of a string, e.g., "3+5/2-(7*9)".
type RPN struct {
	Snum      Stack
	Sop       Stack
	expr      []any
	exprSize  int
	debugMode bool
}

// InitRPN initializes a new RPN calculator with the given size, expression, and capacity.
// It checks if the expression is valid (not too short and has balanced brackets).
// If the expression is valid, it initializes the stacks and converts the expression into an array of any type.
// If the expression is invalid, it returns an error.
// The capacity is used to determine the size of the stacks.
// The size parameter is used to determine the size of the expression.
// The expr parameter is the mathematical expression in string format.
// The function returns a pointer to the RPN struct and an error if any.
func InitRPN(size int, expr string, capacity int) (*RPN, error) {
	if len(expr) < 3 { // 3 is the minimum length for a valid expression (e.g., "1+1")
		return nil, fmt.Errorf("error: invalid expression, expression not enough")
	}
	if !isFullBracket(expr) {
		return nil, fmt.Errorf("error: invalid expression, unbalanced brackets")
	}

	return &RPN{
		Snum:     StackInit(int(capacity / 2)), // Stack
		Sop:      StackInit(int(capacity / 2)), // Stack
		expr:     stringToAnyArray(expr),       // Array
		exprSize: size,                         // Size of the expression
	}, nil
}

// PrintForm converts the expression into Reverse Polish Notation (RPN) format.
// It processes the expression by pushing numbers onto the number stack (Snum) and operators onto the operator stack (Sop).
// When a closing bracket is encountered, it pops all operators until the corresponding opening bracket is found.
// It then pushes the operators back onto the number stack.
// Finally, it combines all items from the operator stack into the number stack and returns the RPN expression as a string.
// The function returns the RPN expression as a string and an error if any.
// The RPN expression is in the form of a string, e.g., "3 5 2 / + 7 9 * -".
// The function also handles operator precedence and ensures that the operators are in the correct order.
func (r *RPN) PrintForm() (result string, err error) {

	if r.debugMode {
		fmt.Println("On Transforming to RPN:")
	}

	// sorting the expression into stack
	for _, v := range r.expr {
		if str, ok := v.(string); ok {
			// if its a number, we push it into the num stack
			num, err := strconv.Atoi(str)
			if err == nil {
				r.Snum.Push(num)

				if r.debugMode {
					fmt.Printf("Pushed to num stack: %v\n", r.Snum.items[:r.Snum.top+1])
				}

				continue
			}

			// if its full bracket, we pop all inside it
			if str == ")" {
				v, err := popInsideBracket(&r.Sop)
				if err != nil {
					return "", err
				}

				// moving inside bracket into num stack
				for _, item := range v {
					r.Snum.Push(item)

					if r.debugMode {
						fmt.Printf("Pushed to num stack from bracket: %v\n", r.Snum.items[:r.Snum.top+1])
					}
				}

				continue
			}

			// if its an operator, we check its priority
			// and pop all operators with higher or equal priority
			for !r.Sop.IsEmpty() {
				v, err := popIfPriority(*r, str)
				if err != nil {
					return "", err
				}
				if v == nil {
					break
				}

				r.Snum.Push(v)

				if r.debugMode {
					fmt.Printf("Popped from Sop to num stack: %v\n", r.Snum.items[:r.Snum.top+1])
				}
			}
			// if its an operator, we push it into the operator stack
			r.Sop.Push(str)

			if r.debugMode {
				fmt.Printf("Pushed to Sop stack: %v\n", r.Sop.items[:r.Sop.top+1])
			}
		}
	}

	// combining it all into one stack for once
	for !r.Sop.IsEmpty() {
		item, err := r.Sop.Pop()
		if err != nil {
			return "", err
		}
		if item == nil {
			continue
		}

		r.Snum.Push(item)

		if r.debugMode {
			fmt.Printf("Pushed to num stack from operand stack: %v\n", r.Snum.items[:r.Snum.top+1])
		}
	}

	reversedResult := ""
	// make all of them into string again
	for !r.Snum.IsEmpty() {
		item, err := r.Snum.Pop()
		if err != nil {
			return "", err
		}

		reversedResult += fmt.Sprintf("%v ", item)
	}

	// reversing the result to get the correct order
	// because we popped from the top of the stack
	// so we need to reverse it to get the correct order
	for i := len(reversedResult) - 1; i >= 0; i-- {
		result += string(reversedResult[i])
	}

	if r.debugMode {
		fmt.Printf("Reversed RPN expression: {%s}\n", reversedResult)
		fmt.Printf("Final RPN expression: {%s}\n", result)
	}

	return result, nil
}

// Calculate evaluates the RPN expression and returns the result as a float64.
// It first converts the RPN expression into an array of any type.
// Then, it iterates through the array, pushing numbers onto the number stack (Snum) and calculating the result for operators.
// It pops two numbers from the number stack, applies the operator, and pushes the result back onto the number stack.
// Finally, it pops the final result from the number stack and returns it as a float64.
// If any error occurs during the calculation, it returns an error.
// The function returns the result as a float64 and an error if any.
// The result is the final calculated value of the RPN expression.
// The RPN expression is expected to be in the form of a string, e.g., "3 5 2 / + 7 9 * -".
func (r *RPN) Calculate() (float64, error) {

	// getting the expression in RPN format
	expr, err := r.PrintForm()
	if err != nil {
		return math.Inf(-1), err
	}

	if r.debugMode {
		fmt.Println("On Calculate:")
	}

	// converting the expression into an array of any type
	exprAny := stringToAnyArray(expr)

	// iterating through the expression array
	for _, v := range exprAny {
		if str, ok := v.(string); ok {
			// if its a number, we push it into the num stack
			num, err := strconv.Atoi(str)
			if err == nil {
				r.Snum.Push(num)

				if r.debugMode {
					fmt.Printf("Pushed to num stack: %v\n", r.Snum.items[:r.Snum.top+1])
				}
			}

			// if its an operator, we pop two numbers from the num stack
			// and apply the operator on them
			// then we push the result back into the num stack
			if op, err := ToOperand(str); err == nil {
				val1, err := r.Snum.Pop()
				if err != nil {
					return math.Inf(-1), err
				}
				val1Float, err := anyToFloat(val1)
				if err != nil {
					return math.Inf(-1), err
				}
				val2, err := r.Snum.Pop()
				if err != nil {
					return math.Inf(-1), err
				}
				val2Float, err := anyToFloat(val2)
				if err != nil {
					return math.Inf(-1), err
				}

				newValue, err := op.Calculate(val2Float, val1Float)
				if err != nil {
					return math.Inf(-1), err
				}

				if r.debugMode {
					fmt.Printf("Calculated: %v %s %v = %v\n", val2Float, op, val1Float, newValue)
				}
				r.Snum.Push(newValue)
				continue
			}
		}
	}

	// popping the final result from the num stack
	result, err := r.Snum.Pop()
	if err != nil {
		return math.Inf(-1), err
	}
	resultFloat, err := anyToFloat(result)
	if err != nil {
		return math.Inf(-1), err
	}

	return resultFloat, nil
}

// stringToAnyArray converts a string into an array of any type.
// It splits the string into individual characters and skips empty strings or spaces.
// It returns an array of any type containing the characters of the string.
// The function is used to convert the expression into an array format for further processing.
// The resulting array can be used for calculations or other operations in the RPN calculator.
func stringToAnyArray(s string) []any {

	// Split the string into individual characters
	aS := strings.Split(s, "")
	var res []any

	// Iterate through the split string and convert each character to any type
	// Skip empty strings or spaces
	for _, v := range aS {
		if v == " " || v == "" || v == "\n" {
			continue
		}

		res = append(res, v)
	}

	return res
}

// anyToFloat converts an any type to a float64.
// It checks the type of the input and converts it accordingly.
// If the input is an int, it converts it to float64.
// If the input is a float64, it returns it as is.
// If the input is a string, it attempts to parse it as a float64.
// If the conversion fails, it returns an error.
// If the input type is not recognized, it returns an error indicating the type cannot be converted.
func anyToFloat(a any) (float64, error) {
	switch v := a.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			return num, nil
		}
		return 0, fmt.Errorf("error: cannot convert string '%s' to float64", v)
	default:
		return 0, fmt.Errorf("error: cannot convert %v (type %T) to float64", a, a)
	}
}

// isFullBracket checks if the given string has balanced brackets.
// It counts the number of left and right brackets and returns true if they are equal.
// If there are no brackets, it also returns true.
// The function is used to validate the expression before processing it in the RPN calculator.
// It ensures that the expression has balanced brackets, which is essential for correct calculations.
func isFullBracket(s string) bool {
	expr := strings.Split(s, "")

	var leftHold, rightHold int
	for _, v := range expr {
		if v == "(" {
			leftHold++
		}
		if v == ")" {
			rightHold++
		}
	}

	return leftHold == rightHold || leftHold == 0 && rightHold == 0
}

// popIfPriority checks if the last operator in the operator stack (Sop) has
// a higher priority than the current operator.
//
// If it does, it pops the last operator from the stack and returns it.
func popIfPriority(r RPN, currStr string) (any, error) {

	// Peek the last operator from the stack
	lastStr, err := r.Sop.Peek()
	if err != nil {
		return nil, err
	}
	if lastStr == nil {
		return nil, nil
	}

	// Convert the last and current operators to Operand types
	lastOp, err := ToOperand(lastStr)
	if err != nil {
		return nil, err
	}
	currOp, err := ToOperand(currStr)
	if err != nil {
		return nil, err
	}

	// Check if the last operator has a higher priority than the current operator
	// If it does, pop the last operator from the stack and return it
	if lastOp.IsPriorityThan(currOp) {
		v, err := r.Sop.Pop()
		if err != nil {
			return nil, err
		}

		return v, nil
	}

	return nil, nil
}

// popInsideBracket pops all items from the operator stack (Sop) until it finds an opening bracket "(".
// It returns the items inside the brackets as a slice of any type.
// If it encounters an error while popping, it returns an error.
// This function is used to handle expressions with brackets in the RPN calculator.
func popInsideBracket(s *Stack) ([]any, error) {
	var items []any

	for !s.IsEmpty() {
		v, err := s.Pop()
		if err != nil {
			return nil, err
		}

		if v == "(" {
			break
		}

		items = append(items, v)
	}

	return items, nil
}

// DebugOn enables debug mode for the RPN calculator.
// In debug mode, additional information is printed to the console during calculations.
// This can help in tracing the steps of the calculation and understanding the flow of the program.
func (r *RPN) DebugOn() {
	r.debugMode = true
	fmt.Println("Debug mode is ON")
}
