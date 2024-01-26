package models

import (
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

func SearchEmployeeByID(employeeID int32) (types.Employee, error) {
	var employee types.Employee
	database, err := Connect()
	if err != nil {
		return employee, err
	}

	err = database.Where("employee_id = ?", employeeID).First(&employee).Error
	return employee, err
}

// func FetchAllEmployees() ([]types.Employee, error) {
// 	db, err := Connection()
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err := db.Query("SELECT employee_id, employee_name, employee_phone, checkpoint_id FROM employees")
// 	db.Close()

// 	if err != nil {
// 		return nil, err
// 	}

// 	var employees []types.Employee
// 	for rows.Next() {
// 		var employee types.Employee
// 		err := rows.Scan(&employee.EmployeeID, &employee.EmployeeName, &employee.EmployeePhone, &employee.CheckpointID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		employees = append(employees, employee)
// 	}

// 	return employees, err
// }
