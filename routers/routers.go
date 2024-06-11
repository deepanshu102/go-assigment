package routers

import (
	"net/http"
	"strconv"

	"github.com/deepanshu102/go-assigment/models"
	"github.com/deepanshu102/go-assigment/repo"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupRouter(repo *repo.EmployeeRepo) *gin.Engine {
	r := gin.Default()
	validate := validator.New()
	r.POST("/employees", func(c *gin.Context) {
		var employee models.Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(employee); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
			return
		}
		createdEmployee := repo.CreateEmployee(employee)
		c.JSON(http.StatusOK, createdEmployee)
	})

	r.GET("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		employee, exists := repo.GetEmployeeByID(id)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
	})

	r.PUT("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var employee models.Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(employee); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
			return
		}

		employee.ID = id
		if !repo.UpdateEmployee(employee) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
	})

	r.DELETE("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if !repo.DeleteEmployee(id) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	r.GET("/employees", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		employees := repo.ListEmployees(page, pageSize)
		c.JSON(http.StatusOK, employees)
	})

	return r
}
