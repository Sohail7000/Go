package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil{
		return 1000 , errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil{
		return 1000, errors.New ("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0622)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic ("Sorry we cannot continue")
	}

	var choice int64
	var depositedAmount float64
	var withdrawAmount float64

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Printf("Please enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is %.2f\n", accountBalance)
		case 2:
			fmt.Println("Enter the amount you wish to deposit: ")
			fmt.Scan(&depositedAmount)
			if depositedAmount < 0 {
				fmt.Println("Deposited amount should be greater than 0")
				continue
			}
			accountBalance += depositedAmount
			fmt.Printf("Your current account balance now is %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Printf("Enter the amount you wish withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("WithdrawAmount amount should be greater than 0")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Sorry you cannot withdraw amount which is more than your accountBalance")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Your current account balance now is %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("GoodBye")
			return
		}
	}
	
}

// if choice == 1 {
// 	fmt.Printf("Your current balance is %.2f\n", accountBalance)
// } else if choice == 2 {
// 	fmt.Println("Enter the amount you wish to deposit: ")
// 	fmt.Scan(&depositedAmount)
// 	if depositedAmount < 0 {
// 		fmt.Println("Deposited amount should be greater than 0")
// 		continue
// 	}
// 	accountBalance += depositedAmount
// 	fmt.Printf("Your current account balance now is %.2f\n", accountBalance)

// } else if choice == 3 {
// 	fmt.Printf("Enter the amount you wish withdraw: ")
// 	fmt.Scan(&withdrawAmount)
// 	if withdrawAmount <= 0 {
// 		fmt.Println("WithdrawAmount amount should be greater than 0")
// 		continue
// 	}
// 	if withdrawAmount > accountBalance {
// 		fmt.Println("Sorry you cannot withdraw amount which is more than your accountBalance")
// 		continue
// 	}
// 	accountBalance -= withdrawAmount
// 	fmt.Printf("Your current account balance now is %.2f\n", accountBalance)
// } else {
// 	fmt.Println("GoodBye")
// 	break
// }