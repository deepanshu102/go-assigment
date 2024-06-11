package db

import (
	"reflect"
	"testing"

	"github.com/deepanshu102/go-assigment/models"
)

func TestCreateEmployee(t *testing.T) {
	store := NewInMemoryStore()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}

	createdEmployee := store.CreateEmployee(employee)
	if createdEmployee.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", createdEmployee.ID)
	}
	if !reflect.DeepEqual(createdEmployee, store.employees[createdEmployee.ID]) {
		t.Errorf("Expected employee %+v, got %+v", createdEmployee, store.employees[createdEmployee.ID])
	}
}

func TestGetEmployeeByID(t *testing.T) {
	store := NewInMemoryStore()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := store.CreateEmployee(employee)

	retrievedEmployee, exists := store.GetEmployeeByID(createdEmployee.ID)
	if !exists {
		t.Error("Expected employee to exist, but it doesn't")
	}
	if !reflect.DeepEqual(createdEmployee, retrievedEmployee) {
		t.Errorf("Expected employee %+v, got %+v", createdEmployee, retrievedEmployee)
	}
}

func TestUpdateEmployee(t *testing.T) {
	store := NewInMemoryStore()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := store.CreateEmployee(employee)

	createdEmployee.Position = "Senior Developer"
	updated := store.UpdateEmployee(createdEmployee)
	if !updated {
		t.Error("Expected employee to be updated, but it wasn't")
	}

	updatedEmployee, exists := store.GetEmployeeByID(createdEmployee.ID)
	if !exists || updatedEmployee.Position != "Senior Developer" {
		t.Errorf("Expected position to be 'Senior Developer', got %s", updatedEmployee.Position)
	}
}

func TestDeleteEmployee(t *testing.T) {
	store := NewInMemoryStore()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := store.CreateEmployee(employee)

	deleted := store.DeleteEmployee(createdEmployee.ID)
	if !deleted {
		t.Error("Expected employee to be deleted, but it wasn't")
	}

	_, exists := store.GetEmployeeByID(createdEmployee.ID)
	if exists {
		t.Error("Expected employee to not exist, but it does")
	}
}

func TestListEmployees(t *testing.T) {
	store := NewInMemoryStore()
	employees := []models.Employee{
		{Name: "John Doe", Position: "Developer", Salary: 60000},
		{Name: "Jane Smith", Position: "Manager", Salary: 80000},
	}

	for _, emp := range employees {
		store.CreateEmployee(emp)
	}

	listedEmployees := store.ListEmployees(1, 1)
	if len(listedEmployees) != 1 {
		t.Errorf("Expected 1 employee, got %d", len(listedEmployees))
	}

	listedEmployees = store.ListEmployees(1, 2)
	if len(listedEmployees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(listedEmployees))
	}
}
