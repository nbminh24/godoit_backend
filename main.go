package main

import (
	"godoit_backend/handlers"
	"godoit_backend/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// Initialize the data store and the handler
	taskStore := store.NewTaskStore()
	taskHandler := handlers.NewTaskHandler(taskStore)

	// API routes
	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetTasks)
		tasks.POST("", taskHandler.CreateTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}

	// Run the server
		router.Run(":8080")
}
