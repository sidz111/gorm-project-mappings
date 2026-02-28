package model

type Employee struct {
	ID      int      `gorm:"primaryKey; autoincrement" json:"id"`
	Name    string   `json:"name"`
	Salary  float64  `json:"salary"`
	Project *Project `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE" json:"project"`
}
