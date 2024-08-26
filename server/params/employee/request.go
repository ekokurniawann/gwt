package params

import (
	"go-web-template/server/models"
	"time"
)


type EmployeeCreate struct {
	NIP     string
	Name    string
	Address string
}


func (e *EmployeeCreate) ParseToModel() *models.Employee {
	employee := models.NewEmployee()
	employee.Address = e.Address
	employee.Name = e.Name
	employee.NIP = e.NIP
	return employee
}


type EmployeeUpdate struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	UpdatedAt time.Time
}


func (e *EmployeeUpdate) ParseToModel() *models.Employee {
	return &models.Employee{
		ID:        e.ID,
		NIP:       e.NIP,
		Name:      e.Name,
		Address:   e.Address,
		UpdatedAt: time.Now(),
	}
}
