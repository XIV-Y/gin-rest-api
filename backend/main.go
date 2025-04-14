package main

import (
	"log"
	"os"

	"github.com/XIV-Y/gin-rest-api/db"
	"github.com/XIV-Y/gin-rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Initialize()
	
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/vnd.api+json")

		c.Next()
	})

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}