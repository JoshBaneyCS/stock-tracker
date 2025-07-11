package main

import (
	"backend/handlers"
	"backend/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Continuing without it.")
	}

	// Create a new Gin router
	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/me", handlers.GetProfile)
		auth.POST("/logout", handlers.Logout)
		// Add more authenticated routes here (e.g. stock routes)
	}

	// Start server on port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
