package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile string = "balance.txt"

func main() {

	account_balance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("Can't continue, sorry.")
	}
	fileops.WriteFloatToFile(account_balance, accountBalanceFile)
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at ", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", account_balance)
		} else if choice == 2 {
			var amt float64
			fmt.Println("Enter amount you want to deposite")
			fmt.Scan(&amt)
			if amt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			account_balance += amt
			fileops.WriteFloatToFile(account_balance, accountBalanceFile)
			fmt.Println("Your updated balance is:", account_balance)
		} else if choice == 3 {
			var amt float64
			fmt.Println("Enter amount you want to withdraw")
			fmt.Scan(&amt)
			if amt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			if amt > account_balance {
				fmt.Println("Withdraw amount cannot be more than current balance")
				continue
			} else {
				account_balance -= amt
				fileops.WriteFloatToFile(account_balance, accountBalanceFile)
				fmt.Println("Your updated balance is:", account_balance)
			}
		} else {
			fmt.Println("Goodbye!")
			break
		}

	}
	fmt.Println("Thanks for choosing our bank")
}
