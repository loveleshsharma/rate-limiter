package rule

import "time"

type Rule struct {
	endpoint        string
	timeUnit        time.Duration
	requestsPerUnit int
}

func NewRule(endpoint string, timeUnit time.Duration, requestsPerUnit int) Rule {
	return Rule{
		endpoint:        endpoint,
		timeUnit:        timeUnit,
		requestsPerUnit: requestsPerUnit,
	}
}
