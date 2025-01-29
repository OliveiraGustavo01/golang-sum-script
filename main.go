package main

import "fmt"

func main() {
	var x, y int
	var op string

	fmt.Print("First number, pleas: ")
	fmt.Scanln(&x)

	fmt.Print("Second number, pleas: ")
	fmt.Scanln(&y)

	fmt.Print("Choose an operation (+, -, *, /): ")
	fmt.Scanln(&op)

	result := calc(x, y, op)
	fmt.Println("Your result is:", result)
}

func calc(x int, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		if y != 0 {
			return x / y
		} else {
			fmt.Println("Error: division by zero")
			return 0
		}
	default:
		fmt.Println("Invalid operation")
		return 0
	}
}
