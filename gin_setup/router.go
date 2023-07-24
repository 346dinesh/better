package gin_setup

import (
	"github.com/346dinesh/better/handlers"
	"github.com/gin-gonic/gin"
)

func RootRouters(route *gin.Engine) {

	// Ping test
	route.GET("/ping", pingHandler)
	// Routes
	route.POST("/login", handlers.LoginHandler)
	route.POST("/dashboard", handlers.DashBaordHandler)

}

func pingHandler(c *gin.Context) {
	c.String(200, "pong")
}
