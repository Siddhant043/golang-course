package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	birthdate string
}

func main() {
	userFirstName := getUserData("Please enter your First Name: ")
	userLastName := getUserData("Please enter your Last Name: ")
	userBirthdate := getUserData("Please enter your Birth Date (MM/DD/YYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
	}
	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(msg string) string {
	fmt.Print(msg)
	var value string
	fmt.Scan(&value)
	return value
}
