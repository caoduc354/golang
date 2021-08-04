package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello")

	fmt.Println(checkNumberNT(9))
}

func checkNumberNT(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; float64(i) <= math.Sqrt(float64(a)); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true

}
