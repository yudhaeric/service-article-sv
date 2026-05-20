package main

import (
	"article-api/app/config"
	"article-api/app/models"
	"article-api/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.ConnectDatabase()

	// Auto Migration
	config.DB.AutoMigrate(&models.Article{})

	// Initialize Router
	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r)

	// Run Server
	r.Run(":8080")
}
