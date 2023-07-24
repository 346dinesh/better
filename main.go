package main

import (
	"fmt"

	"github.com/346dinesh/better/database"
	"github.com/346dinesh/better/gin_setup"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDatabase()
	router := gin.Default()

	router.Use(corsMiddleware)
	gin_setup.RootRouters(router)

	// Start server
	fmt.Println("Server started on http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		fmt.Println("err:", err)
	}
}

// corsMiddleware is a custom middleware to handle CORS headers
func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "*")

	// Handle OPTIONS preflight requests
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
