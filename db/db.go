package db

import (
	"sync"

	"github.com/deepanshu102/go-assigment/models"
)

// InMemoryStore simulates an in-memory database
type InMemoryStore struct {
	mu        sync.Mutex
	employees map[int]models.Employee
	nextID    int
}

// NewInMemoryStore initializes an in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
}

// CreateEmployee adds a new employee to the store
func (store *InMemoryStore) CreateEmployee(employee models.Employee) models.Employee {
	store.mu.Lock()
	defer store.mu.Unlock()

	employee.ID = store.nextID
	store.nextID++
	store.employees[employee.ID] = employee
	return employee
}

// GetEmployeeByID retrieves an employee by ID
func (store *InMemoryStore) GetEmployeeByID(id int) (models.Employee, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()

	employee, exists := store.employees[id]
	return employee, exists
}

// UpdateEmployee updates an existing employee's details
func (store *InMemoryStore) UpdateEmployee(employee models.Employee) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[employee.ID]; !exists {
		return false
	}
	store.employees[employee.ID] = employee
	return true
}

// DeleteEmployee deletes an employee by ID
func (store *InMemoryStore) DeleteEmployee(id int) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[id]; !exists {
		return false
	}
	delete(store.employees, id)
	return true
}

// ListEmployees returns a paginated list of employees
func (store *InMemoryStore) ListEmployees(page, pageSize int) []models.Employee {
	store.mu.Lock()
	defer store.mu.Unlock()

	employees := make([]models.Employee, 0, len(store.employees))
	for _, employee := range store.employees {
		employees = append(employees, employee)
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(employees) {
		return []models.Employee{}
	}

	if end > len(employees) {
		end = len(employees)
	}

	return employees[start:end]
}
