package main

import (
	"fmt"

	"example.com/structs/v2/user"
)

func main() {
	userFirstName := getUserData("Please enter your First Name: ")
	userLastName := getUserData("Please enter your Last Name: ")
	userBirthdate := getUserData("Please enter your Birth Date (MM/DD/YYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserData()
	appUser.OutputUserDetails()

}

func getUserData(msg string) string {
	fmt.Print(msg)
	var value string
	fmt.Scanln(&value) // When using Scanln, hitting enter enter means closing the input.
	return value
}
