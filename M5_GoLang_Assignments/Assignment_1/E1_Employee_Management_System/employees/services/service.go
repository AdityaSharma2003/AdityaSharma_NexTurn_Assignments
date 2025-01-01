package services

import (
	"employee-management/employees/models"
	"errors"
	"fmt"
	"strings"
)

var Employees []models.Employee

func AddEmployee(id int, name string, age int, department string) error {
	for _, employee := range Employees {
		if employee.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	if age <= 18 {
		return errors.New("employee Age must be greater than 18")
	}

	employee := models.Employee{ID: id, Name: name, Age: age, Department: department}

	Employees = append(Employees, employee)

	fmt.Printf("Employee added: %v\n", employee)

	return nil
}

func SearchEmployee(input interface{}) (models.Employee, error) {
	if value, ok := input.(string); ok {
		for _, employee := range Employees {
			if strings.Contains(employee.Name, value) {
				return employee, nil
			}
		}
	} else if value, ok := input.(int); ok {
		for _, employee := range Employees {
			if employee.ID == value {
				return employee, nil
			}
		}
	}
	return models.Employee{}, errors.New("Employee Not Found")
}

func ListEmployeesByDepartment(department string) ([]models.Employee, error) {
	var filteredEmployees []models.Employee
	for _, employee := range Employees {
		if employee.Department == department {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}

	if len(filteredEmployees) == 0 {
		return []models.Employee{}, errors.New("Department Not Found")
	}
	return filteredEmployees, nil
}

func CountEmployeesByDepartment(department string) int {
	count := 0
	for _, employee := range Employees {
		if employee.Department == department {
			count++
		}
	}

	return count
}
