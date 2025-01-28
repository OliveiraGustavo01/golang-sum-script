package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("First number, pleas: ")
	fmt.Scanln(&x)

	fmt.Print("Second number, pleas: ")
	fmt.Scanln(&y)
	soma := calc(x, y)
	fmt.Println("Your sum is", soma)
}

func calc(x int, y int) int {
	return x + y
}
