package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Srinivas1643/zocket_task/backend/config"
	"github.com/Srinivas1643/zocket_task/backend/handlers"
	"github.com/Srinivas1643/zocket_task/backend/middlewares"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Initialize database connection
	db := config.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Apply authentication middleware
	r.Use(middlewares.AuthMiddleware())

	// Authentication routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", handlers.SignUpHandler(db))
		authRoutes.POST("/login", handlers.LoginHandler(db))
	}

	// Task management routes
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("/", handlers.GetTasksHandler(db))
		taskRoutes.POST("/", handlers.CreateTaskHandler(db))
		taskRoutes.PUT("/:id", handlers.UpdateTaskHandler(db))
		taskRoutes.DELETE("/:id", handlers.DeleteTaskHandler(db))
	}

	// AI-powered task recommendation
	r.POST("/tasks/recommend", handlers.AITaskRecommendationHandler())

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if no env variable is set
	}
	log.Println("Server is running on port:", port)
	r.Run(":" + port)
}
