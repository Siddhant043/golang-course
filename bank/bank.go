package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceStr := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceStr), 0644)
}

func main() {

	account_balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("Can't continue, sorry.")
	}
	writeBalanceToFile(account_balance)
	fmt.Println("Welcome to Go Bank!")
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
			writeBalanceToFile(account_balance)
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
				writeBalanceToFile(account_balance)
				fmt.Println("Your updated balance is:", account_balance)
			}
		} else {
			fmt.Println("Goodbye!")
			break
		}

	}
	fmt.Println("Thanks for choosing our bank")
}
