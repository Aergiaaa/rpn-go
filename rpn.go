package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RPN struct {
	Snum     Stack
	Sop      Stack
	expr     []any
	exprSize int
}

func InitRPN(size int, expr string, capacity int) (*RPN, error) {
	if len(expr) < 3 {
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

func (r *RPN) PrintForm() (result string, err error) {

	// sorting the expression into stack
	for _, v := range r.expr {
		if str, ok := v.(string); ok {
			num, err := strconv.Atoi(str)
			if err == nil {
				r.Snum.Push(num)
				continue
			}

			if str == ")" {
				v, err := popInsideBracket(&r.Sop)
				if err != nil {
					return "", err
				}

				// moving inside bracket into num stack
				for _, item := range v {
					r.Snum.Push(item)
				}

				continue
			}

			if !r.Sop.IsEmpty() {
				v, err := popIfPriority(*r, str)
				if err != nil {
					return "", err
				}
				if v != nil {
					r.Snum.Push(v)
					continue
				}
			}

			r.Sop.Push(str)
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

	for i := len(reversedResult) - 1; i >= 0; i-- {
		result += string(reversedResult[i])
	}

	return result, nil
}

func (r *RPN) Calculate() (float64, error) {
	expr, err := r.PrintForm()
	if err != nil {
		return math.Inf(-1), err
	}

	fmt.Printf("Expression in RPN: {%v}\n", expr)

	exprAny := stringToAnyArray(expr)

	for _, v := range exprAny {
		if str, ok := v.(string); ok {
			num, err := strconv.Atoi(str)
			if err == nil {
				r.Snum.Push(num)
			}

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

				newValue, err := op.CalculateString(val2Float, val1Float)
				if err != nil {
					return math.Inf(-1), err
				}

				fmt.Println("Calculated:", val2Float, op, val1Float, "=", newValue) // Debugging line
				r.Snum.Push(newValue)
				continue
			}
		}
	}

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

func stringToAnyArray(s string) []any {
	aS := strings.Split(s, "")
	fmt.Println("Split string into array:", aS) // Debugging line
	var a []any

	for _, v := range aS {
		if v == " " || v == "" || v == "\n" {
			continue // Skip empty strings or spaces
		}

		a = append(a, v)
	}

	fmt.Println("Converted string to any array:", a) // Debugging line

	return a
}

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

func popIfPriority(r RPN, currStr string) (any, error) {
	lastStr, err := r.Sop.Peek()
	if err != nil {
		return nil, err
	}

	if lastStr == nil {
		return nil, nil
	}

	lastOp, err := ToOperand(lastStr)
	if err != nil {
		return nil, err
	}
	currOp, err := ToOperand(currStr)
	if err != nil {
		return nil, err
	}

	if lastOp.IsPriorityThan(currOp) {
		v, err := r.Sop.Pop()
		fmt.Println("Popped from Sop due to priority:", v)                       // Debugging line
		fmt.Println("Sop after pop due to priority:", r.Sop.items[:r.Sop.top+1]) // Debugging line
		fmt.Println("Sop Top: ", r.Sop.top)                                      // Debugging line
		if err != nil {
			return nil, err
		}

		return v, nil
	}

	return nil, nil
}

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
