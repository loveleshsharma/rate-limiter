package pkg

import (
	"github.com/loveleshsharma/rate-limiter/internal/strategy"
	"github.com/loveleshsharma/rate-limiter/pkg/rule"
)

type RateLimiter struct {
	strategy strategy.RateLimitStrategy
}

func (r RateLimiter) IsRequestAllowed() bool {
	return r.strategy.IsRequestAllowed()
}

func NewRateLimiter(strategy strategy.Strategy, rule rule.Rule) RateLimiter {
	return RateLimiter{
		strategy: strategy.GetStrategy(rule),
	}
}
