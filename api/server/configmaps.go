package server

import (
	"net/http"

	"github.com/Josh2604/go-infinite-scroll/api/dependencies"
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine, handlers *dependencies.Handlers) {
	postRoutes(router, handlers)
}

func postRoutes(router *gin.Engine, handlers *dependencies.Handlers) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Running")
	})
	router.POST("/posts", handlers.GetPosts.Handle)
}
