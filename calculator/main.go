package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Calculator functions
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <operator> <num1> <num2>")
		fmt.Println("Operators: add, subtract, multiply, divide")
		return
	}

	// Parse command-line arguments
	operator := os.Args[1]
	num1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid first number.")
		return
	}
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid second number.")
		return
	}

	// Call the appropriate function based on the operator
	switch operator {
	case "add":
		result := Add(num1, num2)
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case "subtract":
		result := Subtract(num1, num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case "multiply":
		result := Multiply(num1, num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case "divide":
		result, err := Divide(num1, num2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Invalid operator. Please use add, subtract, multiply, or divide.")
	}
}
