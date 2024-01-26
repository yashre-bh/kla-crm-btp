package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yashre-bh/kla-crm-btp/pkg/controller"
)

func Start() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")

	employee := api.Group("/employee")
	employee.POST("/add", controller.AddNewEmployee)
	employee.POST("/login", controller.LoginUser)
	// employee.GET("/fetch", controller.FetchAllEmployees)

	checkpoint := api.Group("/checkpoint")
	// checkpoint.GET("/fetch", controller.FetchAllCheckpoints)
	checkpoint.POST("/add", controller.AddCheckpoint)

	fmt.Println("Server listening on :8080...")
	router.Run(":8080")
}
