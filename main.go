package main

import (
	"fmt"
)

func main() {
	fmt.Println("Simple Calculator using Reverse Polish Notation (RPN) (+, -, *, /)")
	fmt.Println("Enter your expression (3+5/2-(4*2)):")

	var expr string
	_, err := fmt.Scanln(&expr)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	rpn, err := InitRPN(len(expr), expr, 100)
	if err != nil {
		fmt.Println("Error initializing RPN:", err)
		return
	}

	result, err := rpn.Calculate()
	if err != nil {
		fmt.Println("Error calculating result:", err)
		return
	}

	fmt.Println("Result:", result)
}
