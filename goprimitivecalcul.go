package main

import (
	"fmt"
	"os"
)

func main() {
	var x int
	var y int
	var operator string

	fmt.Println("Enter first operand")
	fmt.Fscan(os.Stdin, &x)
	fmt.Println("Enter second operand")
	fmt.Fscan(os.Stdin, &y)
	fmt.Println("Select operaton\n+ \n -\n *\n /")
	fmt.Fscan(os.Stdin, &operator)

	switch operator {
	case "+":
		fmt.Println("Result: ", x+y)
	case "-":
		fmt.Println("Result: ", x-y)
	case "*":
		fmt.Println("Result: ", x*y)
	case "/":
		if y != 0 {
			fmt.Println("Result: ", x/y)
		} else {
			fmt.Println("Division by zero error")
		}
	default:
		fmt.Println("Incorrect operation")
	}
}
