package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testfile.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// tim so lon nhat be nhat trung binh
	//thuat toan buble sort

	arr := []int{5, 8, 9, 7, 1, 2, 3}
	// fmt.Println(arr)
	fmt.Println(numberMax(arr, len(arr)))
	fmt.Println(numberMin(arr, len(arr)))
	fmt.Println(numberAvarage(arr, len(arr)))
}
func numberMin(arr []int, n int) int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				v := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = v
			}
		}
	}
	return arr[0]
}

func numberMax(arr []int, n int) int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				v := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = v
			}
		}
	}
	return arr[0]
}

func numberAvarage(arr []int, n int) int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				v := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = v
			}
		}
	}
	return arr[n/2]
}
