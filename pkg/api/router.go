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
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
	}))

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
			employee.DELETE("/delete/:id", c.DeleteEmployee)
			employee.POST("/assign", c.AssignCheckpointToEmployee)
			employee.POST("/change-password/:id", c.ChangeEmployeePassword)
			employee.POST("/purchase-register", c.PurchaseRegister)
			employee.GET("/fetch-all-pending-requests", c.FetchAllPendingRequests)
			employee.POST("/approve-request", c.ApproveByRequestID)
			employee.GET("/fetch-all-resolved-requests", c.FetchAllResolvedRequests)
		}

		checkpoint := admin.Group("/checkpoint")
		{
			checkpoint.POST("/add", c.AddCheckpoint)
			checkpoint.GET("/fetch", c.FetchAllCheckpoints)
			checkpoint.GET("/fetch/:id", c.FetchCheckpointByID)
			checkpoint.DELETE("/delete/:id", c.DeleteCheckpoint)
			checkpoint.GET("fetch/:id/employees", c.GetEmployeesAtCheckpoint)
		}

	}

	worker := api.Group("/worker")
	{
		worker.POST("/checkpoint/incoming-raw-material", m.IsWorkerOrAdminOrSupervisor, c.AddIncomingRawMaterial)
		worker.POST("/checkpoint/post-iqf", m.IsWorkerOrAdminOrSupervisor, c.AddPostIQFRecord)
		worker.POST("/raise-request", m.IsWorkerOrSupervisor, c.RaiseRequest)
		worker.GET("/fetch-pending-requests", m.IsWorkerOrSupervisor, c.FetchPendingRequestsOfEmployee)
		worker.GET("/fetch-resolved-requests", m.IsWorkerOrSupervisor, c.FetchResolvedRequestsOfEmployee)
	}

	supervisor := api.Group("/supervisor")
	{
		supervisor.GET("/pending-forms-to-check", m.IsAdminOrSupervisor, c.PendingFormsToBeCheckedBySupervisor)
		supervisor.GET("/fetch-form-data/:checkpointID/:type/:date", m.IsWorkerOrAdminOrSupervisor, c.FetchFormData)
	}

	fetch := api.Group("/fetch")
	{
		fetch.GET("/incoming-raw-material/all", m.IsWorkerOrAdminOrSupervisor, c.FetchAllIncomingRawMaterialData)
	}

	fmt.Println("Server listening on :8080...")
	router.Run(":8080")
}
