package repo

import (
	"reflect"
	"testing"

	"github.com/deepanshu102/go-assigment/db"
	"github.com/deepanshu102/go-assigment/models"
)

func TestEmployeeRepo(t *testing.T) {
	store := db.NewInMemoryStore()
	repo := NewEmployeeRepo(store)

	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}

	// Test CreateEmployee
	createdEmployee := repo.CreateEmployee(employee)
	if createdEmployee.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", createdEmployee.ID)
	}

	// Test GetEmployeeByID
	retrievedEmployee, exists := repo.GetEmployeeByID(createdEmployee.ID)
	if !exists {
		t.Error("Expected employee to exist, but it doesn't")
	}
	if !reflect.DeepEqual(createdEmployee, retrievedEmployee) {
		t.Errorf("Expected employee %+v, got %+v", createdEmployee, retrievedEmployee)
	}

	// Test UpdateEmployee
	createdEmployee.Position = "Senior Developer"
	updated := repo.UpdateEmployee(createdEmployee)
	if !updated {
		t.Error("Expected employee to be updated, but it wasn't")
	}
	updatedEmployee, exists := repo.GetEmployeeByID(createdEmployee.ID)
	if !exists || updatedEmployee.Position != "Senior Developer" {
		t.Errorf("Expected position to be 'Senior Developer', got %s", updatedEmployee.Position)
	}

	// Test DeleteEmployee
	deleted := repo.DeleteEmployee(createdEmployee.ID)
	if !deleted {
		t.Error("Expected employee to be deleted, but it wasn't")
	}
	_, exists = repo.GetEmployeeByID(createdEmployee.ID)
	if exists {
		t.Error("Expected employee to not exist, but it does")
	}

	// Test ListEmployees
	employees := []models.Employee{
		{Name: "John Doe", Position: "Developer", Salary: 60000},
		{Name: "Jane Smith", Position: "Manager", Salary: 80000},
	}
	for _, emp := range employees {
		repo.CreateEmployee(emp)
	}

	listedEmployees := repo.ListEmployees(1, 1)
	if len(listedEmployees) != 1 {
		t.Errorf("Expected 1 employee, got %d", len(listedEmployees))
	}

	listedEmployees = repo.ListEmployees(1, 2)
	if len(listedEmployees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(listedEmployees))
	}
}
