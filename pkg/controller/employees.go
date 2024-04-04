package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yashre-bh/kla-crm-btp/pkg/middlewares"
	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddNewEmployee(c *gin.Context) {
	var employee types.Employee
	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	employee.DateOfJoining = time.Now()
	password := middlewares.GenerateRandomPassword(10, true, true, true)
	employee.Password = middlewares.HashPassword(password)

	err = models.AddNewEmployee(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add new user",
			"error":   err,
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

func LoginUser(c *gin.Context) {
	var employee types.Employee
	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	data, err := models.FetchPasswordOfEmployee(employee.EmployeeID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Could not find employee with the provided EmployeeID",
			"error":   err,
		})
		return
	}

	if !middlewares.CompareHashedPasswords(employee.Password, data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Incorrect password",
		})
		return
	}
	role, err := models.FetchRoleOfEmployee(employee.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Could not get role of the user",
			"error":   err,
		})
		return
	}

	assignedCheckpoints, err := models.FetchAssignedCheckpoints(employee.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Could not find assigned checkpoints for the user",
			"error":   err,
		})
		return
	}

	token, err := middlewares.CreateJWTClaims(employee.EmployeeID, role, assignedCheckpoints)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Could not Log in user",
			"error":   err,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auth", token, 86400, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully logged in employee",
		"token":   token,
	})
}

func FetchAllEmployees(c *gin.Context) {
	employees, err := models.FetchAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve employees from the database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employees,
	})
}

func FetchEmployeeByID(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid id parameter",
		})
		return
	}

	data, err := models.FetchEmployeeByID(employeeID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "could not retrieve employee data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})

}

func DeleteEmployee(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   "invalid id parameter",
		})
		return
	}

	err = models.DeleteEmployee(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "unable to delete employee",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "employee removed from database",
	})
}

func AssignCheckpointToEmployee(c *gin.Context) {
	var assign types.AssignCheckpoint
	err := c.ShouldBindJSON(&assign)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	_, err = models.FetchCheckpointByID(int(assign.CheckpointID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid checkpoint id",
			"error":   err,
		})
		return
	}

	_, err = models.FetchEmployeeByID(int(assign.EmployeeID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid employee id",
			"error":   err,
		})
		return
	}

	err, isAlreadyAssigned := models.CheckAssignedCheckpoints(&assign)
	if isAlreadyAssigned {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("employee already assigned this checkpoint on %v", assign.AssignedAt),
			"error":   err,
		})
		return
	}

	assign.AssignedAt = time.Now()
	err = models.AssignCheckpointToEmployee(assign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "unable to assign checkpoint to the employee",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "employee assigned checkpoint successfully",
	})
}

func PurchaseRegister(c *gin.Context) {
	var purchase types.PurchaseRegister
	err := c.ShouldBindJSON(&purchase)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	err = models.PurchaseRegister(&purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to register purchase",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "purchase registered successfully",
	})
}

func ChangeEmployeePassword(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid id parameter",
		})
		return
	}

	newPassword := middlewares.GenerateRandomPassword(10, true, true, true)
	hashedPassword := middlewares.HashPassword(newPassword)

	err = models.ChangePasswordOfEmployee(int32(employeeID), hashedPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to change password",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Successfully changed user password",
		"new_password": newPassword,
	})

}

func RaiseRequest(c *gin.Context) {
	var requestRequest types.RaiseRequest
	var requestDBQuery types.RaiseRequestDBQuery
	err := c.ShouldBindJSON(&requestRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	claims, err := middlewares.ExtractJWTClaims(c)

	employeeID, ok := claims["employeeID"].(float64)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to extract employeeID from jwt claims",
			"error":   err,
		})
		return
	}

	requestDBQuery.RequestFrom = int32(employeeID)
	requestDBQuery.RequestDescription = requestRequest.RequestDescription

	err = models.RaisePasswordChangeRequest(&requestDBQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to raise request",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully raised password change request",
	})
}
