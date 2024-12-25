package main

import "fmt"

func main() {

	var account_balance = 1000.0
	fmt.Println("Welcome to Go Bank!")
	for i := 0; i < 2; i++ {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
				return
			}
			account_balance += amt
			fmt.Println("Your updated balance is:", account_balance)
		} else if choice == 3 {
			var amt float64
			fmt.Println("Enter amount you want to withdraw")
			fmt.Scan(&amt)
			if amt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}
			if amt > account_balance {
				fmt.Println("Withdraw amount cannot be more than current balance")
				return
			} else {
				account_balance -= amt
				fmt.Println("Your updated balance is:", account_balance)
			}
		} else {
			fmt.Println("Goodbye!")
		}

		fmt.Println("Your choice is ", choice)
	}

}
