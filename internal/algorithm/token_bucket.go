package algorithm

import (
	"fmt"
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

func NewTokenBucket(capacity int) *TokenBucket {
	tokenBucket := TokenBucket{
		capacity: capacity,
	}

	go startTicker(&tokenBucket)
	return &tokenBucket
}

func startTicker(tokenBucket *TokenBucket) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("resetting bucketCount at: ", t)
			tokenBucket.mx.Lock()
			tokenBucket.bucketCount = 0
			tokenBucket.mx.Unlock()
		}
	}
}
