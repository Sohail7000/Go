package main

import (
	"fmt"
)

func main() {
	// readInput("Revenue: ", &revenue) // pointer based implementation passing address
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	tax_rate := getUserInput("Tax Rate: ")
	ebt, profit, ratio := caculateFinancials(revenue, expenses, tax_rate)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func caculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ((100 - tax_rate) / 100) * ebt
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

// func readInput(inputText string, inputValue *float64) {
// 	fmt.Print(inputText)
// 	fmt.Scan(inputValue)
// }
