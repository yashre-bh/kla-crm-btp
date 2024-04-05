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
	requestDBQuery.Title = requestRequest.Title
	requestDBQuery.RequestDate = time.Now()

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
		"message": "Successfully raised request",
	})
}

func FetchAllPendingRequests(c *gin.Context) {
	pendingRequests, err := models.FetchPendingRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve pending requests from database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pendingRequests,
	})

}

func FetchAllResolvedRequests(c *gin.Context) {
	resolvedRequests, err := models.FetchResolvedRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve resolved requests from database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resolvedRequests,
	})

}

func FetchPendingRequestsOfEmployee(c *gin.Context) {
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

	pendingRequests, err := models.FetchPendingRequestsOfEmployee(int32(employeeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve pending requests from database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pendingRequests,
	})
}

func FetchResolvedRequestsOfEmployee(c *gin.Context) {
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

	resolvedRequests, err := models.FetchResolvedRequestsOfEmployee(int32(employeeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve pending requests from database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resolvedRequests,
	})
}

func ApproveByRequestID(c *gin.Context) {
	var resolveByRequestID types.ResolveByRequestID
	var resolveRequestDBQuery types.ResolveRequestDBQuery
	err := c.ShouldBindJSON(&resolveByRequestID)
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

	resolveRequestDBQuery.RequestID = resolveByRequestID.RequestID
	resolveRequestDBQuery.AcceptedBy = int32(employeeID)
	resolveRequestDBQuery.AdminComment = resolveByRequestID.AdminComment
	resolveRequestDBQuery.Resolved = true
	resolveRequestDBQuery.ResolveDate = time.Now()
	resolveRequestDBQuery.Accepted = resolveByRequestID.Accepted

	err = models.ResolveByRequestID(resolveRequestDBQuery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to resolve request in the db",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "request resolved successfully",
	})
}

func PendingFormsToBeCheckedBySupervisor(c *gin.Context) {
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

	checkpoints, err := models.FetchAllCheckpointsOfEmployee(int32(employeeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve assigned checkpoints of the employee",
			"error":   err,
		})
		return
	}

	var pendingChecksList []types.PendingChecksBySupervisor

	for _, checkpointID := range checkpoints {
		var pendingCheck types.PendingChecksBySupervisor
		checkpoint, err := models.FetchCheckpointByID(int(checkpointID.CheckpointID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to extract checkpoint details",
				"error":   err,
			})
			return
		}
		pendingCheck.Title = checkpoint.CheckpointName
		pendingCheck.Checkpoint = checkpoint.CheckpointID

		// fix when you have more checkpoints

		// switch checkpoint.CheckpointName {
		// case "incoming_raw_material":
		pendingCheckItems, err := models.FetchDataForUncheckedFormsCheckpoint1()
		// }

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to extract pending unchecked items",
				"error":   err,
			})
			return
		}
		pendingCheck.List = *pendingCheckItems
		pendingChecksList = append(pendingChecksList, pendingCheck)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pendingChecksList,
	})

}

func FetchFormData(c *gin.Context) {
	var fetchFormDataRequest types.FetchFormDataRequest
	err := c.ShouldBindJSON(&fetchFormDataRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err,
		})
		return
	}

	// add switch case stmt when more tables made
	data, err := models.FetchFormDataFromCheckpoint1(fetchFormDataRequest.CheckpointID, fetchFormDataRequest.BatchCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve form data specified",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
