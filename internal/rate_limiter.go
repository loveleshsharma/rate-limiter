package internal

import "github.com/loveleshsharma/rate-limiter/internal/algorithm"

type RateLimiter struct {
	algo algorithm.RateLimitAlgorithm
}

func (r RateLimiter) IsRequestAllowed() bool {
	return r.algo.IsRequestAllowed()
}

func NewRateLimiter(algo algorithm.RateLimitAlgorithm) RateLimiter {
	return RateLimiter{
		algo: algo,
	}
}
