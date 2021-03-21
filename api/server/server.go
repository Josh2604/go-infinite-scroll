package server

import (
	"github.com/Josh2604/go-infinite-scroll/api/dependencies"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func Start() {
	router := gin.New()

	handlers := dependencies.Exec()

	router.Use(cors.Default())
	routes(router, handlers)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
