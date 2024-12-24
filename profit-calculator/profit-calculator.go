package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")

	expences := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calcFinancials(revenue, expences, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calcFinancials(revenue, expences, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expences

	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	return ebt, profit, ratio
}
