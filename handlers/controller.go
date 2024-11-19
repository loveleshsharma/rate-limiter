package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/loveleshsharma/rate-limiter/pkg"
	"net/http"
)

type Controller struct {
	rateLimiter pkg.RateLimiter
}

func (c Controller) RateLimit(ctx *gin.Context) {

	isAllowed := c.rateLimiter.IsRequestAllowed()

	if isAllowed == true {
		ctx.Status(http.StatusOK)
		return
	}

	ctx.Status(http.StatusTooManyRequests)
}

func NewController(rateLimiter pkg.RateLimiter) Controller {
	return Controller{
		rateLimiter: rateLimiter,
	}
}
