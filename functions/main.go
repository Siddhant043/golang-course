package main

import "fmt"

type funcInt func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubleNums := transformNumbers(&numbers, double)
	fmt.Println(doubleNums)
}

func transformNumbers(nums *[]int, transform funcInt) []int {
	dNums := []int{}

	for _, val := range *nums {
		dNums = append(dNums, transform(val))
	}
	return dNums
}

func double(num int) int {
	return num * 2
}
