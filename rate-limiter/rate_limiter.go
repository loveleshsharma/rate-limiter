package rate_limiter

import "github.com/loveleshsharma/rate-limiter/rate-limiter/algorithm"

type RateLimiter struct {
	algo algorithm.RateLimitAlgorithm
}

func NewRateLimiter() RateLimiter {
	return RateLimiter{}
}
