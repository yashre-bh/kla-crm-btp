package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	fmt.Println(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
	}

	if Roles[claims["role"]] != types.ADMIN {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
	}

	c.Next()
}

func IsSupervisor(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	fmt.Println(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
	}

	if Roles[claims["role"]] != types.SUPERVISOR {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
	}

	c.Next()
}

func IsWorker(c *gin.Context) {
	claims, err := ExtractJWTClaims(c)
	fmt.Println(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "failed to extract jwt claims",
			"error":   err,
		})
	}

	if Roles[claims["role"]] != types.WORKER {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "user not authorised for this action",
		})
	}

	c.Next()
}
