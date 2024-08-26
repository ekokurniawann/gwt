package models

import "time"


type Employee struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}


func NewEmployee() *Employee {
	return &Employee{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}


func (e *Employee) GetCreatedAtString() string {
	return e.CreatedAt.Format(time.RFC3339)
}

func (e *Employee) GetUpdatedAtString() string {
	return e.UpdatedAt.Format(time.RFC3339)
}
