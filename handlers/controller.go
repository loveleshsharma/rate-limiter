package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/internal"
	"net/http"
)

type Controller struct {
	rateLimiter internal.RateLimiter
}

func (c Controller) RateLimit(ctx *gin.Context) {

	isAllowed := c.rateLimiter.IsRequestAllowed()

	if isAllowed == true {
		ctx.Status(http.StatusOK)
		return
	}

	ctx.Status(http.StatusTooManyRequests)
}

func NewController(rateLimiter internal.RateLimiter) Controller {
	return Controller{
		rateLimiter: rateLimiter,
	}
}
