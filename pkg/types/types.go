package types

import (
	"time"
)

type Role string

const (
	ADMIN      Role = "ADMIN"
	SUPERVISOR Role = "SUPERVISOR"
	WORKER     Role = "WORKER"
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

type CheckpointID struct {
	CheckpointID int32 `json:"checkpoint_id"`
}

type JWTClaims struct {
	EmployeeID int32
	Role       Role
}

type AssignCheckpoint struct {
	EmployeeID   int32     `json:"employee_id"`
	CheckpointID int32     `json:"checkpoint_id"`
	AssignedAt   time.Time `json:"assigned_at"`
}

type RaiseRequest struct {
	// RequestID          int32  `json:"request_id"`
	// RequestFrom        int    `json:"request_from"`
	RequestDescription string `json:"request_description"`
	// Accepted           bool   `json:"accepted"`
	// AcceptedBy         int    `json:"accepted_by"`
	// AdminComment       string `json:"admin_comment"`
}

type RaiseRequestDBQuery struct {
	RequestFrom        int32     `json:"request_from"`
	RequestDescription string    `json:"request_description"`
	RequestDate        time.Time `json:"request_date"`
}

type PendingRequests struct {
	RequestID          int32     `json:"request_id"`
	RequestFrom        int       `json:"request_from"`
	RequestDescription string    `json:"request_description"`
	RequestDate        time.Time `json:"request_date"`
}

type ResolveByRequestID struct {
	RequestID    int32  `json:"request_id"`
	Accepted     bool   `json:"accepted"`
	AdminComment string `json:"admin_comment"`
}

type ResolveRequestDBQuery struct {
	RequestID    int32     `json:"request_id"`
	Accepted     bool      `json:"accepted"`
	Resolved     bool      `json:"resolved"`
	AdminComment string    `json:"admin_comment"`
	AcceptedBy   int32     `json:"accepted_by"`
	ResolveDate  time.Time `json:"resolve_date"`
}

type ResolvedRequests struct {
	RequestID          int32     `json:"request_id"`
	RequestFrom        int       `json:"request_from"`
	RequestDescription string    `json:"request_description"`
	RequestDate        time.Time `json:"request_date"`
	Accepted           bool      `json:"accepted"`
	Resolved           bool      `json:"resolved"`
	AdminComment       string    `json:"admin_comment"`
	AcceptedBy         int32     `json:"accepted_by"`
	ResolveDate        time.Time `json:"resolve_date"`
}

type PendingChecksBySupervisor struct {
	Title      string              `json:"title"`
	Checkpoint int32               `json:"checkpoint"`
	List       []PendingCheckItems `json:"pending_check_items"`
}

type PendingCheckItems struct {
	Name            string    `json:"name"`
	DateOfArrival   time.Time `json:"date_of_arrival"`
	AddedByEmployee int32     `json:"added_by_employee"`
	BatchCode       string    `json:"batch_code"`
}
