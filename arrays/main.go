package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{3.2, 3.3, 3.4, 3.5}

	fmt.Println(prices)
	productNames[2] = "Something"

	fmt.Println(productNames)
	fmt.Println(prices[2])
}
