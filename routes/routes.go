package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/handlers"
	"github.com/loveleshsharma/rate-limiter/internal/strategy"
	"github.com/loveleshsharma/rate-limiter/pkg"
	"github.com/loveleshsharma/rate-limiter/pkg/rule"
	"time"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	rateLimiter := pkg.NewRateLimiter(strategy.Tokenbucket, rule.NewRule(time.Second, 5))
	rateLimitController := handlers.NewController(rateLimiter)

	router.GET("/ratelimit", rateLimitController.RateLimit)

	return router
}
