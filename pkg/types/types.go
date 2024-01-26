package types

import (
	"time"
)

type Role string

const (
	ADMIN   Role = "ADMIN"
	MANAGER Role = "MANAGER"
	WORKER  Role = "WORKER"
)

type Employee struct {
	EmployeeID    int32     `json:"employee_id" gorm:"primaryKey;autoIncrement;not null"`
	Password      string    `json:"password" gorm:"not null"`
	Name          string    `json:"name" gorm:"not null"`
	DateOfBirth   time.Time `json:"dob" gorm:"not null"`
	DateOfJoining time.Time `json:"doj" gorm:"not null"`
	Designation   string    `json:"designation" gorm:"not null"`
	Department    string    `json:"department"`
	Address       string    `json:"address" gorm:"not null"`
	Phone         string    `json:"phone" gorm:"unique;not null"`
	Email         string    `json:"email" gorm:"unique"`
	Role          Role      `json:"role" gorm:"not null"`
}

type Checkpoint struct {
	CheckpointID   int32  `json:"checkpoint_id" gorm:"primaryKey;autoIncrement"`
	CheckpointName string `json:"checkpoint_name" gorm:"not null;unique"`
}
