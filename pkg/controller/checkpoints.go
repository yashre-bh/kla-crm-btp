package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddCheckpoint(c *gin.Context) {
	var checkpoint types.Checkpoint
	err := c.ShouldBindJSON(&checkpoint)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	err = models.AddCheckpoint(&checkpoint)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add new checkpoint",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":       true,
		"message":       "Checkpoint successfully added to database",
		"checkpoint_id": checkpoint.CheckpointID,
	})
}

func FetchAllCheckpoints(c *gin.Context) {
	checkpoints, err := models.FetchAllCheckpoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve checkpoints from the database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    checkpoints,
	})

}

func FetchCheckpointByID(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "invalid id parameter",
		})
		return
	}

	data, err := models.FetchCheckpointByID(employeeID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "could not retrieve checkpoint data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func DeleteCheckpoint(c *gin.Context) {
	checkpointID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "invalid id parameter",
		})
		return
	}

	err = models.DeleteCheckpoint(checkpointID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "could not delete checkpoint",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "checkpoint removed",
	})
}

func GetEmployeesAtCheckpoint(c *gin.Context) {
	var employees []types.Employee
	checkpointID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "invalid id parameter",
			"error":   err,
		})
	}

	err = models.GetEmployeesAtCheckpoint(checkpointID, &employees)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "could not fetch employees at the checkpoint",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employees,
	})

}
