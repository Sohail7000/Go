package main

import (
	"errors"
	"fmt"
	"os"
)

const financesFile = "finances.txt"

func main() {
	// readInput("Revenue: ", &revenue) // pointer based implementation passing address
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	tax_rate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := caculateFinancials(revenue, expenses, tax_rate)
	financesReport := fmt.Sprintf("ebt: %.1f\nprofit: %.1f\nratio: %.1f\n", ebt, profit, ratio)
	storeInputInFile(financesReport)
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

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value must be a postive number")
		// panic("User input can be zero or a negative number")
	}
	return userInput, nil
}

func storeInputInFile(financesReport string) {
	os.WriteFile(financesFile, []byte(financesReport), 0622)
}

// func readInput(inputText string, inputValue *float64) {
// 	fmt.Print(inputText)
// 	fmt.Scan(inputValue)
// }
