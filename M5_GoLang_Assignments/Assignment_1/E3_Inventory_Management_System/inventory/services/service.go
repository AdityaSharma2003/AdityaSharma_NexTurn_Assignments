package services

import (
	"errors"
	"fmt"
	"inventory-management/inventory/models"
	"strconv"
	"strings"
)

var Inventory []models.Product

func AddProduct(id int, name string, price float64, stock int) error {
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	for _, product := range Inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}

	newProduct := models.Product{ID: id, Name: name, Price: price, Stock: stock}
	Inventory = append(Inventory, newProduct)
	fmt.Println("Product added successfully!")
	return nil
}

func UpdateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range Inventory {
		if product.ID == id {
			Inventory[i].Stock = stock
			fmt.Println("Stock updated successfully!")
			return nil
		}
	}
	return errors.New("product with ID not found")
}

func SearchProduct(input string) (models.Product, error) {
	id, err := strconv.Atoi(input)

	if err != nil {
		for _, item := range Inventory {
			if strings.Contains(item.Name, input) {
				return item, nil
			}
		}
	} else {
		for _, item := range Inventory {
			if item.ID == id {
				return item, nil
			}
		}
	}
	return models.Product{}, errors.New("product not found")
}

func DisplayInventory() {
	fmt.Printf("\n%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range Inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
	fmt.Println()
}

func SortInventoryBy(option string) error {
	var temp models.Product
	n := len(Inventory)
	switch strings.ToLower(option) {
	case "price":
		for i := n - 1; i >= 1; i-- {
			for j := 0; j <= i-1; j++ {
				if Inventory[j].Price > Inventory[j+1].Price {
					temp = Inventory[j]
					Inventory[j] = Inventory[j+1]
					Inventory[j+1] = temp
				}
			}
		}
		fmt.Println("Products sorted by price.")
	case "stock":
		for i := n - 1; i >= 1; i-- {
			for j := 0; j <= i-1; j++ {
				if Inventory[j].Stock > Inventory[j+1].Stock {
					temp = Inventory[j]
					Inventory[j] = Inventory[j+1]
					Inventory[j+1] = temp
				}
			}
		}
		fmt.Println("Products sorted by stock.")
	default:
		return errors.New("invalid sort option (use 'price' or 'stock')")
	}
	return nil
}
