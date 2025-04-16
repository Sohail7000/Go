package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetValueFromFile(accountBalanceFile)

	if err != nil {
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

		fmt.Println("We are available throughout the world even at: ", randomdata.Country(randomdata.FullCountry))
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
