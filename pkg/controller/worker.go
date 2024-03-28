package controller

import (
	// "fmt"
	"net/http"
	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"

	"github.com/yashre-bh/kla-crm-btp/pkg/middlewares"
	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	// "github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddIncomingRawMaterial(c *gin.Context) {
	var incomingRawMaterial types.IncomingRawMaterial
	err := c.ShouldBindJSON(&incomingRawMaterial)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}
	today := time.Now().Format("02-01-06")
	batchCode, err := middlewares.CreateBatchCode(incomingRawMaterial.Name, today)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to assign batch code",
			"error":   err,
		})
		return
	}

	incomingRawMaterial.LotNumber = batchCode
	err = models.AddIncomingRawMaterial(&incomingRawMaterial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add new incoming raw material",
			"error":   err,
		})
		return
	}

	err = models.AddToActiveBatches(batchCode, today, incomingRawMaterial.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add to active batches",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "new raw material data entry made",
		"batch_code": batchCode,
	})
}
