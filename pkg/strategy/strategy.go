package strategy

import (
	"github.com/loveleshsharma/rate-limiter/internal/strategies"
	"github.com/loveleshsharma/rate-limiter/pkg/rule"
)

type Strategy int

const (
	Tokenbucket Strategy = iota
)

func (s Strategy) GetStrategy(rule rule.Rule) RateLimitStrategy {
	var rateLimitStrategy RateLimitStrategy
	switch s {
	case Tokenbucket:
		rateLimitStrategy = strategies.NewTokenBucket(rule)
	default:
		rateLimitStrategy = strategies.NewTokenBucket(rule)
	}

	return rateLimitStrategy
}

type RateLimitStrategy interface {
	IsRequestAllowed() bool
}
