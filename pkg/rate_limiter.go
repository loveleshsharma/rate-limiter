package pkg

import "github.com/loveleshsharma/rate-limiter/pkg/strategy"

const (
	tokenBucketDefaultCapacity = 5
)

type RateLimiter struct {
	strategy strategy.RateLimitStrategy
}

func (r RateLimiter) IsRequestAllowed() bool {
	return r.strategy.IsRequestAllowed()
}

func NewRateLimiter(rateLimitType strategy.RateLimitType) RateLimiter {
	rateLimiter := RateLimiter{}
	switch rateLimitType {
	case strategy.TokenBucketStrategy:
		rateLimiter.strategy = strategy.NewTokenBucket(tokenBucketDefaultCapacity)
	}
	return rateLimiter
}
