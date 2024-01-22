package types

type Employee struct {
	EmployeeID    int32  `json:"employee_id"`
	EmployeeName  string `json:"employee_name"`
	EmployeePhone string `json:"employee_phone"`
	CheckpointID  int32  `json:"checkpoint_id"`
}
