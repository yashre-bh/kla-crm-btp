package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddNewEmployee(c *gin.Context) {
	var employee types.Employee
	err := c.ShouldBindJSON(&employee)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request payload",
		})
		return
	}

	employee.DateOfJoining = time.Now()
	password := GenerateRandomPassword(10, true, true, true)
	employee.Password = HashPassword(password)

	err = models.AddNewEmployee(employee)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to add new user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":            true,
		"message":            "User successfully added to database",
		"employee_id":        employee.EmployeeID,
		"temporary_password": password,
	})
}

// func FetchAllEmployees(c *gin.Context) {
// 	employees, err := models.FetchAllEmployees()
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"error":   "Failed to retrieve employees from the database",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"data":    employees,
// 	})
// }
