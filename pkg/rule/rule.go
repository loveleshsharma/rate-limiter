package rule

import "time"

type Rule struct {
	TimeUnit        time.Duration
	RequestsPerUnit int
}

func NewRule(timeUnit time.Duration, requestsPerUnit int) Rule {
	return Rule{
		TimeUnit:        timeUnit,
		RequestsPerUnit: requestsPerUnit,
	}
}
