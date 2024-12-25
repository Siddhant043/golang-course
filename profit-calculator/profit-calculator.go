package main

import (
	"errors"
	"fmt"
	"os"
)

func checkValidInput(num float64) error {
	if num <= 0 {
		return errors.New("Please enter a value larger than 0")
	}
	return nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue := getUserInput("Revenue: ")

	expences := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calcFinancials(revenue, expences, taxRate)

	fmt.Printf("%0.2f\n", ebt)
	fmt.Printf("%0.2f\n", profit)
	fmt.Printf("%0.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	err := checkValidInput(userInput)
	if err != nil {
		panic(err)
	}
	return userInput
}

func calcFinancials(revenue, expences, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expences
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
