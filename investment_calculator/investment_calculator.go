package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")

	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value : %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Formatted Future Value : %1f \n", realFutureValue)
	fmt.Print(formattedFV, formattedRFV)
	// fmt.Printf("Future Value : %.1f\n Future Value(adjusted for inflation) : %.1f \n", futureValue , realFutureValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
