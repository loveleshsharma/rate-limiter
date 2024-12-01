package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RateLimitMiddleware(rateLimiter RateLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isAllowed := rateLimiter.IsRequestAllowed(); isAllowed == false {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		ctx.Next()
	}
}
