package servies

import (
	"errors"
	"fmt"
	params "go-web-template/server/params/employee"
	repositories "go-web-template/server/repositories/employee"
)

type EmployeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}


func (s *EmployeeService) CreateEmployee(param *params.EmployeeCreate) (*params.EmployeeSingleView, error) {
	
	if err := validateEmployeeCreate(param); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	
	employee := param.ParseToModel()

	
	if err := s.repo.Save(*employee); err != nil {
		return nil, fmt.Errorf("failed to save employee: %w", err)
	}

	return &params.EmployeeSingleView{
		ID:        employee.ID, 
		NIP:       employee.NIP,
		Name:      employee.Name,
		Address:   employee.Address,
		CreatedAt: employee.GetCreatedAtString(),
		UpdatedAt: employee.GetUpdatedAtString(),
	}, nil
}


func (s *EmployeeService) GetAllEmployees() (map[string]params.EmployeeSingleView, error) {
	employees, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve employees: %w", err)
	}

	employeeViews := make(map[string]params.EmployeeSingleView)
	for _, employee := range employees {
		employeeViews[employee.NIP] = params.EmployeeSingleView{
			ID:        employee.ID,
			NIP:       employee.NIP,
			Name:      employee.Name,
			Address:   employee.Address,
			CreatedAt: employee.GetCreatedAtString(),
			UpdatedAt: employee.GetUpdatedAtString(),
		}
	}

	return employeeViews, nil
}


func (s *EmployeeService) GetEmployeeByID(id int) (*params.EmployeeSingleView, error) {
	employee, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve employee with ID %d: %w", id, err)
	}

	if employee.ID == 0 {
		return nil, errors.New("employee not found")
	}

	return &params.EmployeeSingleView{
		ID:        employee.ID,
		NIP:       employee.NIP,
		Name:      employee.Name,
		Address:   employee.Address,
		CreatedAt: employee.GetCreatedAtString(),
		UpdatedAt: employee.GetUpdatedAtString(),
	}, nil
}


func (s *EmployeeService) UpdateEmployeeByID(param *params.EmployeeUpdate) (*params.EmployeeSingleView, error) {
	
	if err := validateEmployeeUpdate(param); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	
	employee, err := s.repo.FindByID(param.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find employee: %w", err)
	}

	if employee.ID == 0 {
		return nil, errors.New("employee not found")
	}

	
	updatedEmployee := param.ParseToModel()
	updatedEmployee.CreatedAt = employee.CreatedAt 

	if err := s.repo.UpdateByID(*updatedEmployee); err != nil {
		return nil, fmt.Errorf("failed to update employee: %w", err)
	}

	return &params.EmployeeSingleView{
		ID:        updatedEmployee.ID,
		NIP:       updatedEmployee.NIP,
		Name:      updatedEmployee.Name,
		Address:   updatedEmployee.Address,
		CreatedAt: updatedEmployee.GetCreatedAtString(),
		UpdatedAt: updatedEmployee.GetUpdatedAtString(),
	}, nil
}


func (s *EmployeeService) DeleteEmployeeByID(id int) error {
	
	err := s.repo.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("failed to delete employee with ID %d: %w", id, err)
	}

	return nil
}


func validateEmployeeCreate(param *params.EmployeeCreate) error {
	if param.NIP == "" {
		return errors.New("NIP is required")
	}
	if param.Name == "" {
		return errors.New("Name is required")
	}
	if param.Address == "" {
		return errors.New("Address is required")
	}
	
	return nil
}


func validateEmployeeUpdate(param *params.EmployeeUpdate) error {
	if param.ID <= 0 {
		return errors.New("Invalid ID")
	}
	if param.NIP == "" {
		return errors.New("NIP is required")
	}
	if param.Name == "" {
		return errors.New("Name is required")
	}
	if param.Address == "" {
		return errors.New("Address is required")
	}

	return nil
}
