package main
import (
	"fmt"
)

func main(){
	var revenue float64
	var expenses float64
	var tax_rate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&tax_rate)
	
	EBT := revenue - expenses
	profit := ((100 - tax_rate) /100 ) * EBT
	ratio := EBT/profit
	fmt.Println(EBT) 
	fmt.Println(profit)
	fmt.Println(ratio)


}