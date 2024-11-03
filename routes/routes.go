package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/handlers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	rateLimitController := handlers.NewController()

	router.GET("/ratelimit", rateLimitController.RateLimit)

	return router
}
