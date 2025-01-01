package employees

import (
	"employee-management/employees/models"
	"employee-management/employees/services"
)

func AddEmployee(id int, name string, age int, department string) error {
	return services.AddEmployee(id, name, age, department)
}

func SearchEmployee(input interface{}) (models.Employee, error) {
	return services.SearchEmployee(input)
}

func ListEmployeesByDepartment(department string) ([]models.Employee, error) {
	return services.ListEmployeesByDepartment(department)
}

func CountEmployeesByDepartment(department string) int {
	return services.CountEmployeesByDepartment(department)
}
