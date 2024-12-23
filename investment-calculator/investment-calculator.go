package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)
	formatedRealValueStr := fmt.Sprintf("Future value adjusted after inflation: %.2f\n", futureRealValue)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf(formatedRealValueStr)

}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	return fv
}

func calculateFutureRealValue(futureValue, inflationRate, years float64) float64 {
	frv := futureValue / math.Pow(1+inflationRate/100, years)
	return frv
}
