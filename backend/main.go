package main

import (
	"github.com/atakanargn/task-mgmt/backend/database"
	"github.com/atakanargn/task-mgmt/backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// CORS middleware'i
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Geliştirme için tüm kaynaklara izin ver
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.GET("/tasks", handlers.GetTasks)
		api.POST("/tasks", handlers.CreateTask)
		api.PATCH("/tasks/:id", handlers.UpdateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	router.Run(":8080")
}
