package main

import (
	"fmt"
	"example.com/structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	adminUser := user.NewAdmin("sohai@gmail.com", "hello")
	if err!=nil{
		fmt.Println(err)
		return
	}
	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// func outputUserDetails (u *user){
// 	// pointers to struct are an exception and go allows it without dereferencing it
// 	// normally you gotta derefernce it// go allows it
// 	fmt.Println((*u).firstName, u.lastName, u.birthDate)
// }
