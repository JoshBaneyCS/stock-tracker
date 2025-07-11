package main

import (
	"log"
	"os"

	"handlers"
	"backend/middleware"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment vars.")
	}

	// Connect to DB
	models.ConnectDatabase()

	// Init router
	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/me", handlers.GetProfile)
		auth.POST("/logout", handlers.Logout)

		auth.GET("/favorites", handlers.GetFavorites)
		auth.POST("/favorites", handlers.SetFavorites)

		auth.GET("/stocks", handlers.GetStockData)
		auth.GET("/stock-history/:symbol", handlers.GetStockHistory)

		auth.GET("/settings", handlers.GetSettings)
		auth.POST("/settings", handlers.UpdateSettings)

		auth.POST("/alerts", handlers.CreateAlert)
		auth.GET("/alerts", handlers.GetAlerts)
		auth.DELETE("/alerts/:id", handlers.DeleteAlert)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
