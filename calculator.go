package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Println(num1 / num2)
		}
	default:
		fmt.Println("Invalid operator")
	}

}
