package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
}

func (u User) OutputUserDetails() {
	// this is a shortcut provided by go, technically it is (*u).firstName), you have to dereference it normally
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// to mutate data in a struct as shown in below method you should pass the pointer so actual struct is mutated rather than its copy
func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
}

// This is how we write a constructor for a struct, you can also use pointer approach to avoid creating a copy variable
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
	}, nil
}
