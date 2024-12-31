package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//slicesExample()
	//dynamicArrayExample()
	arrayPractice()
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

func arrayPractice() {
	hobbies := [3]string{"Coding", "Riding", "Gaming"}
	// 1
	fmt.Println("Hobbies list: ", hobbies)

	// 2
	fmt.Println("First Element: ", hobbies[0])
	fmt.Println("Second and Third Element: ", hobbies[1:3])

	// 3
	firstTwoElements := hobbies[:2]
	fmt.Println("First and Second Element: ", hobbies[:2])
	fmt.Println("First and Second Element: ", hobbies[0:2])

	// 4

	firstTwoElements = firstTwoElements[1:3]
	fmt.Println("Second and Last Element: ", firstTwoElements)

	//5

	goalsArr := []string{}

	goalsArr = append(goalsArr, "learning")
	goalsArr = append(goalsArr, "personalGrowth")
	fmt.Println("Goals array: ", goalsArr)

	// 6
	goalsArr[1] = "personalExperience"
	goalsArr = append(goalsArr, "careerGrowth")
	fmt.Println("Updated Goals array: ", goalsArr)

	// 7
	productsArr := []Product{{id: "phone1", title: "Phone 1", price: 22.60}, {id: "phone2", title: "Phone 2", price: 40.60}}

	fmt.Println("Products: ", productsArr)

	newProduct := Product{id: "phone3", title: "Phone 3", price: 99.10}
	productsArr = append(productsArr, newProduct)
	fmt.Println("Updated Products: ", productsArr)
}
