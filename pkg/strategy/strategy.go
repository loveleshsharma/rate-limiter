package strategy

type RateLimitStrategy interface {
	IsRequestAllowed() bool
}

type RateLimitType int

const (
	TokenBucketStrategy RateLimitType = iota
)
