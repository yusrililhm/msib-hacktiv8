package repository

import (
	"database/sql"
	"fmt"
	"go-postgresql/model"
)

type EmployeesRepository struct {
	*sql.DB
}

type EmployeesRepositoryImpl interface {
	CreateEmployee(employees model.Employees) (model.Employees, error)
	GetEmployee() ([]model.Employees, error)
	UpdateEmployee(employee model.Employees) (string, error)
	DeleteEmployee(id int) (string, error)
}

func NewEmployeesRepository(db *sql.DB) EmployeesRepositoryImpl {
	return &EmployeesRepository{
		DB: db,
	}
}

// CreateEmployee implements EmployeesRepositoryImpl.
func (r *EmployeesRepository) CreateEmployee(employees model.Employees) (model.Employees, error) {
	employee := model.Employees{}
	sql := `INSERT INTO employee (full_name, email, age, division) VALUES ($1, $2, $3, $4) RETURNING *;`

	if err := r.QueryRow(sql, employees.FullName, employees.Email, employees.Age, employees.Division).Scan(&employee.Id, &employee.FullName, &employee.Email, &employee.Age, &employee.Division); err != nil {
		return employee, err
	}

	defer r.Close()
	return employee, nil
}

// GetEmployee implements EmployeesRepositoryImpl.
func (r *EmployeesRepository) GetEmployee() ([]model.Employees, error) {
	employees := []model.Employees{}

	sql := "SELECT * FROM employee"
	row, err := r.Query(sql)

	if err != nil {
		return employees, err
	}

	defer row.Close()

	for row.Next() {
		employee := model.Employees{}

		if err := row.Scan(&employee.Id, &employee.FullName, &employee.Email, &employee.Age, &employee.Division); err != nil {
			return employees, nil
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

// UpdateEmployee implements EmployeesRepositoryImpl.
func (r *EmployeesRepository) UpdateEmployee(employee model.Employees) (string, error) {
	sql := "UPDATE employee SET full_name = $2, email = $3, age = $4, division = $5 WHERE id = $1"
	_, err := r.Exec(sql, employee.Id, employee.FullName, employee.Email, employee.Age, employee.Division)

	if err != nil {
		return fmt.Sprintf("Can't update employee : employee id %d not found", employee.Id), err
	}

	defer r.Close()
	return fmt.Sprintf("Update employee id %d success", employee.Id), nil
}

// DeleteEmployee implements EmployeesRepositoryImpl.
func (r *EmployeesRepository) DeleteEmployee(id int) (string, error) {
	sql := "DELETE FROM employee WHERE id = $1"
	_, err := r.Exec(sql, id)

	if err != nil {
		return fmt.Sprintf("Can't delete employee : employee id %d not found", id), err
	}

	return fmt.Sprintf("delete employee id %d success", id), nil
}