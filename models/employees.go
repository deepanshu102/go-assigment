package models

// Employee represents an employee record
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Position string  `json:"position" validate:"required"`
	Salary   float64 `json:"salary" validate:"required,gt=0"`
}
