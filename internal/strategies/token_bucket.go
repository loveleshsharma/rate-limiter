package strategies

import (
	"fmt"
	"github.com/loveleshsharma/rate-limiter/pkg/rule"
	"sync"
	"time"
)

type TokenBucket struct {
	mx sync.Mutex

	bucketCount int
	capacity    int
}

func (b *TokenBucket) IsRequestAllowed() bool {
	b.mx.Lock()
	defer b.mx.Unlock()

	if b.bucketCount < b.capacity {
		b.bucketCount++
		return true
	}

	return false
}

func NewTokenBucket(rule rule.Rule) *TokenBucket {
	tokenBucket := TokenBucket{
		capacity: rule.RequestsPerUnit,
	}

	go tokenBucket.startTicker(rule)
	return &tokenBucket
}

func (b *TokenBucket) startTicker(rule rule.Rule) {
	ticker := time.NewTicker(time.Duration(rule.RequestsPerUnit) * rule.TimeUnit)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("resetting bucketCount at: ", t)
			b.mx.Lock()
			b.bucketCount = 0
			b.mx.Unlock()
		}
	}
}
