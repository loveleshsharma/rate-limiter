package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/handlers"
	"github.com/loveleshsharma/rate-limiter/pkg"
	"github.com/loveleshsharma/rate-limiter/pkg/strategy"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	rateLimiter := pkg.NewRateLimiter(strategy.TokenBucketStrategy)
	rateLimitController := handlers.NewController(rateLimiter)

	router.GET("/ratelimit", rateLimitController.RateLimit)

	return router
}
