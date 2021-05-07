package server

import (
	"github.com/kisstupid/go-gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.Authenticate())

	router.GET("/")

	return router
}
