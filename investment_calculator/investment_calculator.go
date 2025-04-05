package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value : %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Formatted Future Value : %1f \n", realFutureValue)
	fmt.Print(formattedFV,formattedRFV )
	// fmt.Printf("Future Value : %.1f\n Future Value(adjusted for inflation) : %.1f \n", futureValue , realFutureValue)


}
g