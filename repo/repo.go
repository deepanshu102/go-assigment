package repo

import (
	"github.com/deepanshu102/go-assigment/db"
	"github.com/deepanshu102/go-assigment/models"
)

type EmployeeRepo struct {
	store *db.InMemoryStore
}

func NewEmployeeRepo(store *db.InMemoryStore) *EmployeeRepo {
	return &EmployeeRepo{store: store}
}

func (repo *EmployeeRepo) CreateEmployee(employee models.Employee) models.Employee {
	return repo.store.CreateEmployee(employee)
}

func (repo *EmployeeRepo) GetEmployeeByID(id int) (models.Employee, bool) {
	return repo.store.GetEmployeeByID(id)
}

func (repo *EmployeeRepo) UpdateEmployee(employee models.Employee) bool {
	return repo.store.UpdateEmployee(employee)
}

func (repo *EmployeeRepo) DeleteEmployee(id int) bool {
	return repo.store.DeleteEmployee(id)
}

func (repo *EmployeeRepo) ListEmployees(page, pageSize int) []models.Employee {
	return repo.store.ListEmployees(page, pageSize)
}
