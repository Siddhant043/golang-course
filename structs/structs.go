package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	birthdate string
}

func (u user) outputUserDetails() {
	// this is a shortcut provided by go, technically it is (*u).firstName), you have to dereference it normally
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// to mutate data in a struct as shown in below method you should pass the pointer so actual struct is mutated rather than its copy
func (u *user) clearUserData() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your First Name: ")
	userLastName := getUserData("Please enter your Last Name: ")
	userBirthdate := getUserData("Please enter your Birth Date (MM/DD/YYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
	}
	appUser.outputUserDetails()
	appUser.clearUserData()
	appUser.outputUserDetails()
}

func getUserData(msg string) string {
	fmt.Print(msg)
	var value string
	fmt.Scan(&value)
	return value
}
