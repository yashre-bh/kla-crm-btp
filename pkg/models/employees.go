package models

import (
	"fmt"

	"github.com/yashre-bh/kla-crm-btp/pkg/types"
	"gorm.io/gorm"
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

func DeleteEmployee(employeeID int) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Delete(&types.Employee{}, employeeID).Error

	return err
}

func AssignCheckpointToEmployee(assign types.AssignCheckpoint) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("employee_checkpoint").Create(&assign).Error
	return err
}

func CheckAssignedCheckpoints(assign *types.AssignCheckpoint) (error, bool) {
	database, err := Connect()
	if err != nil {
		return err, false
	}

	err = database.Table("employee_checkpoint").Where("checkpoint_id = ? AND employee_id = ?", assign.CheckpointID, assign.EmployeeID).First(&assign).Error

	if err == gorm.ErrRecordNotFound {
		return err, false
	}

	return err, true
}

func PurchaseRegister(purchase *types.PurchaseRegister) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("purchase_register").Create(&purchase).Error
	return err
}
