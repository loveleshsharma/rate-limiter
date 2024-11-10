package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/handlers"
	"github.com/loveleshsharma/rate-limiter/internal"
	"github.com/loveleshsharma/rate-limiter/internal/strategy"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	rateLimiter := internal.NewRateLimiter(strategy.NewTokenBucket(5))
	rateLimitController := handlers.NewController(rateLimiter)

	router.GET("/ratelimit", rateLimitController.RateLimit)

	return router
}
