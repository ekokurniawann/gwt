package repositories

import (
	"database/sql"
	"fmt"
	"go-web-template/server/models"
)

type EmployeeRepository interface {
	Save(employee models.Employee) error
	FindAll() (map[int]models.Employee, error)
	FindByID(id int) (models.Employee, error)
	UpdateByID(employee models.Employee) error
	DeleteByID(id int) error
}

type employeeRepo struct {
	DB *sql.DB
}


func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{
		DB: db,
	}
}


func (e *employeeRepo) Save(employee models.Employee) error {

	query := `
	INSERT INTO employees (
		nip, name, address, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id
	`

	err := e.DB.QueryRow(query, employee.NIP, employee.Name, employee.Address, employee.CreatedAt, employee.UpdatedAt).Scan(&employee.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}


func (e *employeeRepo) FindAll() (map[int]models.Employee, error) {

	query := `
	SELECT
		id, nip, name, address, created_at, updated_at
	FROM
		employees
	`


	rows, err := e.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close() 

	
	employees := make(map[int]models.Employee)

	
	for rows.Next() {
		var employee models.Employee

		
		if err := rows.Scan(
			&employee.ID, &employee.NIP, &employee.Name, &employee.Address,
			&employee.CreatedAt, &employee.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		
		employees[employee.ID] = employee
	}

	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return employees, nil
}


func (e *employeeRepo) FindByID(id int) (models.Employee, error) {
	
	query := `
	SELECT
		id, nip, name, address, created_at, updated_at
	FROM
		employees
	WHERE
		id = $1
	`

	
	var employee models.Employee

	
	err := e.DB.QueryRow(query, id).Scan(
		&employee.ID, &employee.NIP, &employee.Name, &employee.Address,
		&employee.CreatedAt, &employee.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		
		return models.Employee{}, nil
	} else if err != nil {
		
		return models.Employee{}, fmt.Errorf("failed to retrieve employee with ID %d: %w", id, err)
	}

	
	return employee, nil
}


func (e *employeeRepo) UpdateByID(employee models.Employee) error {
	
	query := `
	UPDATE employees 
	SET name=$1, address=$2, nip=$3, updated_at=$4
	WHERE id=$5
	`


	result, err := e.DB.Exec(query, employee.Name, employee.Address, employee.NIP, employee.UpdatedAt, employee.ID)
	if err != nil {
		
		return fmt.Errorf("failed to update employee with ID %d: %w", employee.ID, err)
	}

	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not verify update of employee with ID %d: %w", employee.ID, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no employees found with ID %d", employee.ID)
	}

	
	return nil
}


func (e *employeeRepo) DeleteByID(id int) error {
	
	query := `
	DELETE FROM employees
	WHERE id = $1
	`


	result, err := e.DB.Exec(query, id)
	if err != nil {
		
		return fmt.Errorf("failed to delete employee with ID %d: %w", id, err)
	}

	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		
		return fmt.Errorf("could not verify deletion of employee with ID %d: %w", id, err)
	}

	
	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with ID %d", id)
	}

	return nil
}
