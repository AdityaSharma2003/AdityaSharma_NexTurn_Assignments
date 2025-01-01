package users

import "bank-transaction/Users/services"

func AddAccount(id int, name string, balance float64) {
	services.AddAccount(id, name, balance)
}

func Deposit(id int, amount float64) error {
	return services.Deposit(id, amount)
}

func Withdraw(id int, amount float64) error {
	return services.Withdraw(id, amount)
}

func ViewBalance(id int) error {
	return services.ViewBalance(id)
}
