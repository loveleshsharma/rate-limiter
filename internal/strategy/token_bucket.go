package strategy

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

	go tokenBucket.startTicker()
	return &tokenBucket
}

func (b *TokenBucket) startTicker() {
	ticker := time.NewTicker(time.Second * 5)

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
