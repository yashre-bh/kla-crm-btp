package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddCheckpoint(c *gin.Context) {
	var checkpoint types.Checkpoint
	err := c.ShouldBindJSON(&checkpoint)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request payload",
		})
		return
	}

	err = models.AddCheckpoint(&checkpoint)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to add new checkpoint",
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
