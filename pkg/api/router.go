package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yashre-bh/kla-crm-btp/pkg/controller"
)

func Start() {
	router := gin.Default()

	api := router.Group("/api")

	employee := api.Group("/employee")
	employee.GET("/fetch", controller.FetchAllEmployees)
	employee.POST("/add", controller.AddEmployee)

	checkpoint := api.Group("/checkpoint")
	checkpoint.GET("/fetch", controller.FetchAllCheckpoints)
	checkpoint.POST("/add", controller.AddCheckpoint)

	fmt.Println("Server listening on :8080...")
	router.Run(":8080")
}
