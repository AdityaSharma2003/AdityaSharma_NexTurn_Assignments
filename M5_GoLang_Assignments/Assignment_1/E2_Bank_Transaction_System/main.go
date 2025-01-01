package main

import (
	"bank-transaction/Users/services"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("\n--- Bank Transaction Menu ---")
		fmt.Println("1. Add Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Balance")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			var id int
			var name string
			var balance float64
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Name: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error occured in reading name. Error message : ", err)
				name = "John Doe"
			} else {
				name = strings.TrimSpace(input)
			}
			fmt.Print("Enter Initial Balance: ")
			fmt.Scanln(&balance)
			services.AddAccount(id, name, balance)

		case 2:
			var id int
			var amount float64
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scanln(&amount)
			err := services.Deposit(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case 3:
			var id int
			var amount float64
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Withdraw Amount: ")
			fmt.Scanln(&amount)
			err := services.Withdraw(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case 4:
			var id int
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			err := services.ViewBalance(id)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println()

		case 5:
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}
