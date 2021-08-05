package main

import "fmt"

func main() {

	fmt.Println(calculator(6.2, "+", 2.0))
	fmt.Println(calculator(6.2, "-", 2))
	fmt.Println(calculator(6.2, "*", 2))
	fmt.Println(calculator(6.2, "/", 2))
	fmt.Println(calculator(6.2, "/", 0))

}

func calculator(a float64, c string, b float64) float64 {
	var result float64

L:

	switch c {
	case "+":
		result = a + b

	case "-":
		result = a - b

	case "*":
		result = a * b

	case "/":
		if b == 0 {
			fmt.Println("Can not divide by zero")
			break L
		} else {
			result = a / b

		}
	default:
		fmt.Println("chuoi khong xacs dinh")
	}

	return result
}
