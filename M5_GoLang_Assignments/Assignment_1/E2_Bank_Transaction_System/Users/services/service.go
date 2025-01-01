package services

import (
	"bank-transaction/Users/models"
	"errors"
	"fmt"
	"time"
)

var Users []models.Users

func AddAccount(id int, name string, balance float64) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	initialHistory := []string{fmt.Sprintf("Account created on %s", currentTime)}
	user := models.Users{ID: id, Name: name, Balance: balance, History: initialHistory}
	Users = append(Users, user)
	fmt.Printf("User Account added: %v\n", user)
}

func Deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be greater than zero")
	}
	for i, user := range Users {
		if user.ID == id {
			Users[i].Balance += amount
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			msg := fmt.Sprintf("Deposited amount (%s) for ID(%v): %.2f", currentTime, id, amount)
			Users[i].History = append(Users[i].History, msg)
			return nil
		}
	}
	return errors.New("user with the ID not found")
}

func Withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be greater than zero")
	}
	for i, user := range Users {
		if user.ID == id {
			Users[i].Balance -= amount
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			msg := fmt.Sprintf("Withdrawn amount (%s) for ID(%v): %.2f", currentTime, id, amount)
			Users[i].History = append(Users[i].History, msg)
			return nil
		}
	}
	return errors.New("user with the ID not found")
}

func ViewBalance(id int) error {
	for _, user := range Users {
		if user.ID == id {
			fmt.Printf("Current Balance of user with ID (%v) is: %v", id, user.Balance)
			return nil
		}
	}
	return errors.New("user with the ID not found")
}
