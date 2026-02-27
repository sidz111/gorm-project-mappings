package model

import (
	"time"
)

type Project struct {
	ID         int       `gorm:"primaryKey;autoincrement" json:"id"`
	Title      string    `json:"title"`
	AssignDate time.Time `json:"assignDate"`
	EmployeeID int       `json:"employeeId"`
}
