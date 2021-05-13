package main

import (
	"os"

	"github.com/devrodriguez/muevete-fitness-go-api/cmd/api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	server := server.New()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	// Run server
	server.Run(":" + port)
}
