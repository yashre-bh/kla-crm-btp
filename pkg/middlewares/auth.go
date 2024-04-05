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

func IsWorkerOrAdminOrSupervisor(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
		return
	}

	if Roles[claims["role"]] == types.WORKER || Roles[claims["role"]] == types.SUPERVISOR || Roles[claims["role"]] == types.ADMIN {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})

	}
}
