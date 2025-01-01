package main

import (
	"employee-management/employees"
	"employee-management/employees/models"
	"fmt"
)

func main() {
	var filteredEmployees []models.Employee
	var employee models.Employee
	var err error

	err = employees.AddEmployee(1, "aditya sharma", 22, "IT")
	if err != nil {
		fmt.Println("Error message: ", err)
	}
	err = employees.AddEmployee(2, "virat singh", 25, "IT")
	if err != nil {
		fmt.Println("Error message: ", err)
	}
	err = employees.AddEmployee(3, "sanjeevani arora", 21, "HR")
	if err != nil {
		fmt.Println("Error message: ", err)
	}

	employee, err = employees.SearchEmployee("sanjeevani")
	if err != nil {
		fmt.Println("Error message: ", err)
	}
	fmt.Println("Searched Employee: ", employee)

	employee, err = employees.SearchEmployee(1)
	if err != nil {
		fmt.Println("Error message: ", err)
	}
	fmt.Println("Searched Employee: ", employee)

	filteredEmployees, err = employees.ListEmployeesByDepartment("IT")
	if err != nil {
		fmt.Println("Error message: ", err)
	}
	fmt.Println("Filtered Employees: ", filteredEmployees)

	cnt := employees.CountEmployeesByDepartment("IT")
	fmt.Print("The total count of employees in IT department are: ", cnt)
}
