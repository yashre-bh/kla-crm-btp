package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	c "github.com/yashre-bh/kla-crm-btp/pkg/controller"
	m "github.com/yashre-bh/kla-crm-btp/pkg/middlewares"
)

func Start() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")

	api.POST("/login", c.LoginUser)

	admin := api.Group("/admin")
	admin.Use(m.IsAdmin)
	{
		employee := admin.Group("/employee")
		{
			employee.POST("/add", c.AddNewEmployee)
			employee.GET("/fetch", c.FetchAllEmployees)
			employee.GET("/fetch/:id", c.FetchEmployeeByID)
		}

		checkpoint := admin.Group("/checkpoint")
		{
			checkpoint.POST("/add", c.AddCheckpoint)
			checkpoint.GET("/fetch", c.FetchAllCheckpoints)
		}

	}

	fmt.Println("Server listening on :8080...")
	router.Run(":8080")
}
