package algorithm

type RateLimitAlgorithm interface {
	IsRequestAllowed() bool
}
