package main

import (
	"fmt"

	"github.com/346dinesh/better/gin_setup"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	gin_setup.RootRouters(router)

	// Start server
	fmt.Println("Server started on http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		fmt.Println("err:", err)
	}
}
