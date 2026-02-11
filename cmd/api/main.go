package main

import (
	"log"
	"os"

	_ "gousers/docs"
	"gousers/internal/database"
	"gousers/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Users API
// @version 1.0
// @description This is a sample server for a user management API.
// @host localhost:8084
// @BasePath /api/v1
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	database.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Serve Static Landing Page
	r.StaticFile("/", "./web/index.html")

	// API Routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", handlers.GetUsers)
		v1.POST("/users", handlers.CreateUser)
		v1.GET("/users/:id", handlers.GetUser)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
