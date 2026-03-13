package main

import "fmt"

func main() {
	inputNumber1 := []int{1, 3, 2, 4}
	inputNumber2 := []int{1, 4, 8, 9}
	inputNumber3 := []int{9, 9, 9, 9}
	fmt.Println(plusOne(inputNumber1))
	fmt.Println(plusOne(inputNumber2))
	fmt.Println(plusOne(inputNumber3))
}

func plusOne(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < 9 {
			arr[i]++
			return arr
		}
		arr[i] = 0
	}
	// return arr
	return append([]int{1}, arr...)
}
