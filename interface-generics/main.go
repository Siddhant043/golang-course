package main

import "fmt"

func main() {
	a := "fruit: "
	b := "mango"

	add(a, b)
}

func add[T int | float64 | string](a, b T) {
	fmt.Println(a + b)
}
