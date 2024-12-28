package main

import (
	"errors"
	"fmt"
)

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

// This is how we write a constructor for a struct, you can also use pointer approach to avoid creating a copy variable
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your First Name: ")
	userLastName := getUserData("Please enter your Last Name: ")
	userBirthdate := getUserData("Please enter your Birth Date (MM/DD/YYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserData()
	appUser.outputUserDetails()

}

func getUserData(msg string) string {
	fmt.Print(msg)
	var value string
	fmt.Scanln(&value) // When using Scanln, hitting enter enter means closing the input.
	return value
}
