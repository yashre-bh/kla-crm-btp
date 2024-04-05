package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

var (
	Roles = map[interface{}]types.Role{
		"ADMIN":      types.ADMIN,
		"SUPERVISOR": types.SUPERVISOR,
		"WORKER":     types.WORKER,
	}
)

func IsAdmin(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
		return
	}

	if Roles[claims["role"]] != types.ADMIN {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
		return
	}

	c.Next()
}

func IsSupervisor(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
		return
	}

	if Roles[claims["role"]] != types.SUPERVISOR {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
		return
	}

	c.Next()
}

func IsWorker(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
		return
	}

	if Roles[claims["role"]] != types.WORKER {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
		return
	}

	c.Next()
}

func IsWorkerOrSupervisor(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
		return
	}
	if Roles[claims["role"]] != types.WORKER && Roles[claims["role"]] != types.SUPERVISOR {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
		return
	}
	c.Next()
}

// func IsEmployeeAssignedToCheckpoint(c *gin.Context) {
// 	claims, err := ExtractJWTClaims(c)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"success": false,
// 			"message": "failed to extract jwt claims",
// 			"error":   err,
// 		})
// 	}

// 	if Roles[claims["role"]] == types.ADMIN {
// 		c.Next()
// 	} else if Roles[claims["role"]] == types.SUPERVISOR || Roles[claims["role"]] == types.WORKER {
// 		checkpointID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"success": false,
// 				"message": "specified checkpoint does not exist",
// 				"error":   err,
// 			})
// 		}
// 		if claims["checkpoint"] == nil {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"success": false,
// 				"message": "employee not assigned to any checkpoint",
// 			})
// 		}
// 		for _, checkpoint := range claims["checkpoint"].([]interface{}) {
// 			if num, ok := checkpoint.(float64); ok {
// 				if int32(num) == int32(checkpointID) {
// 					c.Next()
// 					return
// 				}
// 			} else {
// 				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 					"success": false,
// 					"message": "invalid checkpoint value in claim",
// 				})
// 				return
// 			}
// 		}

// 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "employee not assigned to this checkpoint",
// 		})
// 	} else {
// 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "user not authorised for this action",
// 		})
// 	}
// }

func IsEmployeeAssignedToCheckpoint(employeeID int32, checkpointID int32) (bool, error) {
	data, err := models.FetchAssignedCheckpoints(employeeID)
	if err != nil {
		return false, errors.New("failed to retrieve assigned checkpoints")
	}

	var assigned bool = false
	for _, val := range data {
		if id, ok := val.(int32); ok {
			if !ok {
				return false, errors.New("failed to retrieve assigned checkpoints (1)")
			}

			if id == checkpointID {
				assigned = true
				break
			}
		}
	}

	return assigned, nil
}
