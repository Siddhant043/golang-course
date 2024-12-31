package main

import "fmt"

func main() {
	//slicesExample()
	dynamicArrayExample()
}

func slicesExample() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{3.2, 3.3, 3.4, 3.5}

	fmt.Println(prices)
	productNames[2] = "Something"

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[:1]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}

func dynamicArrayExample() {
	prices := []float64{10.99, 8.90}
	fmt.Println(prices)

	prices = append(prices, 5.99)
	fmt.Println(prices)
}
