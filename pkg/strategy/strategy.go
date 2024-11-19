package strategy

type RateLimitStrategy interface {
	IsRequestAllowed() bool
}
