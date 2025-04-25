package main

import (
	"fmt"
	"time"
)


func main() {

	age := 26
	var agePointer *int
	agePointer = &age

	fmt.Println("Age", *agePointer)

	editAgeToAdultYears(agePointer)

	fmt.Println(age)
}

func editAgeToAdultYears(age *int)  {
	*age = *age - 18
}
