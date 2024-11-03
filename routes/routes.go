package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	return router
}
