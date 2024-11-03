package main

import (
	"fmt"
	"github.com/loveleshsharma/rate-limiter/internal"
	"github.com/loveleshsharma/rate-limiter/internal/algorithm"
)

func main() {

	rtLimiter := internal.NewRateLimiter(algorithm.NewTokenBucket(3))
	fmt.Println(rtLimiter.IsRequestAllowed())

}
