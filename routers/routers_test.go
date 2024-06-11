package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepanshu102/go-assigment/db"
	"github.com/deepanshu102/go-assigment/models"
	"github.com/deepanshu102/go-assigment/repo"
	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	store := db.NewInMemoryStore()
	employeeRepo := repo.NewEmployeeRepo(store)
	return SetupRouter(employeeRepo)
}

func TestCreateEmployeeRoute(t *testing.T) {
	router := setupTestRouter()

	// Valid employee
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	jsonValue, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	// Invalid employee (missing name)
	invalidEmployee := models.Employee{Position: "Developer", Salary: 60000}
	jsonValue, _ = json.Marshal(invalidEmployee)
	req, _ = http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code 400, got %d", w.Code)
	}

	// Invalid employee (salary less than or equal to 0)
	invalidEmployee = models.Employee{Name: "John Doe", Position: "Developer", Salary: 0}
	jsonValue, _ = json.Marshal(invalidEmployee)
	req, _ = http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code 400, got %d", w.Code)
	}
}

func TestGetEmployeeByIDRoute(t *testing.T) {
	router := setupTestRouter()

	// Create an employee first
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	jsonValue, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now get the employee by ID
	req, _ = http.NewRequest("GET", "/employees/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	var retrievedEmployee models.Employee
	json.Unmarshal(w.Body.Bytes(), &retrievedEmployee)

	if retrievedEmployee.ID != 1 {
		t.Errorf("Expected employee ID to be 1, got %d", retrievedEmployee.ID)
	}
	if retrievedEmployee.Name != "John Doe" {
		t.Errorf("Expected employee name to be 'John Doe', got %s", retrievedEmployee.Name)
	}
}

func TestUpdateEmployeeRoute(t *testing.T) {
	router := setupTestRouter()

	// Create an employee first
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	jsonValue, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now update the employee
	employee.Position = "Senior Developer"
	jsonValue, _ = json.Marshal(employee)
	req, _ = http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	var updatedEmployee models.Employee
	json.Unmarshal(w.Body.Bytes(), &updatedEmployee)

	if updatedEmployee.Position != "Senior Developer" {
		t.Errorf("Expected employee position to be 'Senior Developer', got %s", updatedEmployee.Position)
	}
}

func TestDeleteEmployeeRoute(t *testing.T) {
	router := setupTestRouter()

	// Create an employee first
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	jsonValue, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now delete the employee
	req, _ = http.NewRequest("DELETE", "/employees/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Expected status code 204, got %d", w.Code)
	}

	// Try to get the deleted employee
	req, _ = http.NewRequest("GET", "/employees/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %d", w.Code)
	}
}
