package pkg

import "github.com/loveleshsharma/rate-limiter/pkg/strategy"

type RateLimiter struct {
	strategy strategy.RateLimitStrategy
}

func (r RateLimiter) IsRequestAllowed() bool {
	return r.strategy.IsRequestAllowed()
}

func NewRateLimiter(strategy strategy.RateLimitStrategy) RateLimiter {
	return RateLimiter{
		strategy: strategy,
	}
}
