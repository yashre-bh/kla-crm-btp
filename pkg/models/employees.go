package models

import (
	"fmt"

	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddNewEmployee(employee *types.Employee) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Create(&employee).Error
	return err

}

func FetchPasswordOfEmployee(employeeID int32) (types.Employee, error) {
	var employee types.Employee
	database, err := Connect()
	if err != nil {
		return employee, err
	}

	err = database.Select("password").Where("employee_id = ?", employeeID).First(&employee).Error
	return employee, err
}

func FetchAllEmployees() ([]types.Employee, error) {
	database, err := Connect()
	if err != nil {
		return nil, err
	}

	var employees []types.Employee
	err = database.Omit("password").Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, err
}

func FetchEmployeeByID(employeeID int) (types.Employee, error) {
	var employee types.Employee
	database, err := Connect()
	if err != nil {
		return employee, err
	}

	err = database.Omit("password").Where("employee_id = ?", employeeID).First(&employee).Error

	fmt.Println(employee)
	return employee, err
}
