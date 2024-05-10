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

	checkpoint, err := models.FetchCheckpointByName("incoming_raw_material")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "couldn't fetch checkpoint data",
			"error":   err,
		})
		return
	}

	isAssigned, err := middlewares.IsEmployeeAssignedToCheckpoint(int32(employeeID), checkpoint.CheckpointID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "couldn't check whether employee is assigned to the checkpoint",
			"error":   err,
		})
		return
	}

	if !isAssigned {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not assigned to this checkpoint",
		})
		return
	}

	var incomingRawMaterial types.IncomingRawMaterial
	err = c.ShouldBindJSON(&incomingRawMaterial)

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

	err = models.AddToMasterTracking(batchCode, &incomingRawMaterial.DateOfArrival)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add to batch tracking",
			"error":   err,
		})
		return
	}

	incomingRawMaterial.BatchCode = batchCode

	var incomingRawMaterialDBQuery types.IncomingRawMaterialDBQuery
	incomingRawMaterialDBQuery.Name = incomingRawMaterial.Name
	incomingRawMaterialDBQuery.DateOfArrival = incomingRawMaterial.DateOfArrival
	incomingRawMaterialDBQuery.VehicleNumber = incomingRawMaterial.VehicleNumber
	incomingRawMaterialDBQuery.BatchCode = incomingRawMaterial.BatchCode
	incomingRawMaterialDBQuery.Variety = incomingRawMaterial.Variety
	incomingRawMaterialDBQuery.ReceivedFrom = incomingRawMaterial.ReceivedFrom
	incomingRawMaterialDBQuery.Supplier = incomingRawMaterial.Supplier
	incomingRawMaterialDBQuery.WeightSupplier = incomingRawMaterial.WeightSupplier
	incomingRawMaterialDBQuery.WeightWM = incomingRawMaterial.WeightWM
	incomingRawMaterialDBQuery.Rate = incomingRawMaterial.Rate
	incomingRawMaterialDBQuery.Color = incomingRawMaterial.Color
	incomingRawMaterialDBQuery.Texture = incomingRawMaterial.Texture
	incomingRawMaterialDBQuery.Size = incomingRawMaterial.Size
	incomingRawMaterialDBQuery.Maturity = incomingRawMaterial.Maturity
	incomingRawMaterialDBQuery.Aroma = incomingRawMaterial.Aroma
	incomingRawMaterialDBQuery.Appearance = incomingRawMaterial.Appearance
	incomingRawMaterialDBQuery.WeightAccepted = incomingRawMaterial.WeightAccepted
	incomingRawMaterialDBQuery.WeighmentSlipNumber = incomingRawMaterial.WeighmentSlipNumber
	incomingRawMaterialDBQuery.QuantityRejected = incomingRawMaterial.QuantityRejected
	incomingRawMaterialDBQuery.Remarks = incomingRawMaterial.Remarks
	incomingRawMaterialDBQuery.AddedByEmployee = int32(employeeID)

	err = models.AddIncomingRawMaterial(&incomingRawMaterialDBQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add new incoming raw material",
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

func AddPostIQFRecord(c *gin.Context) {
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

	var postIQF types.PostIQF
	err = c.ShouldBindJSON(&postIQF)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	err = models.BatchProgressToCheckpoint2(postIQF.BatchCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update master tracking",
			"error":   err,
		})
		return
	}

	var postIQFDBQuery types.PostIQFDBQuery

	postIQFDBQuery.BatchCode = postIQF.BatchCode
	postIQFDBQuery.BlancherBeltSpeed = postIQF.BlancherBeltSpeed
	postIQFDBQuery.BlancherTemperature = postIQF.BlancherTemperature
	postIQFDBQuery.CoolerBeltSpeed = postIQF.CoolerBeltSpeed
	postIQFDBQuery.CoolerTemperature = postIQF.CoolerTemperature
	postIQFDBQuery.AddedByEmployee = int32(employeeID)
	postIQFDBQuery.SprayNozzleBlancher = postIQF.SprayNozzleBlancher
	postIQFDBQuery.SprayNozzleWasher = postIQF.SprayNozzleWasher
	postIQFDBQuery.SprayNozzleCooler = postIQF.SprayNozzleCooler
	postIQFDBQuery.SprayNozzlePrecooler = postIQF.SprayNozzlePrecooler
	postIQFDBQuery.SprayNozzleBeltSpeed1 = postIQF.SprayNozzleBeltSpeed1
	postIQFDBQuery.SprayNozzleBeltSpeed2 = postIQF.SprayNozzleBeltSpeed2
	postIQFDBQuery.IQFProductTemperature = postIQF.IQFProductTemperature
	postIQFDBQuery.IQFAirTemperature = postIQF.IQFAirTemperature
	postIQFDBQuery.IQFCoilTemperature = postIQF.IQFCoilTemperature
	postIQFDBQuery.BagNumber = postIQF.BagNumber
	postIQFDBQuery.TotalBag = postIQF.TotalBag
	postIQFDBQuery.DateAdded = postIQF.DateAdded

	err = models.AddPostIQFRecord(&postIQFDBQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add Post IQF record",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post IQF record added successfully",
	})
}
