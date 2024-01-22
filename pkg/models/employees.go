package models

import (
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddEmployee(employee types.Employee) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	insertSql := "INSERT INTO employees (employee_name, employee_phone, checkpoint_id) VALUES (?, ?, ?)"
	_, err = db.Exec(insertSql, employee.EmployeeName, employee.EmployeePhone, employee.CheckpointID)
	db.Close()

	return err
}

func FetchAllEmployees() ([]types.Employee, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT employee_id, employee_name, employee_phone, checkpoint_id FROM employees")
	db.Close()

	if err != nil {
		return nil, err
	}

	var employees []types.Employee
	for rows.Next() {
		var employee types.Employee
		err := rows.Scan(&employee.EmployeeID, &employee.EmployeeName, &employee.EmployeePhone, &employee.CheckpointID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, err
}
