package middleware

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	cond     *sync.Cond
	attempts map[string]int
	max      int
}

func NewRateLimiter(max int) *RateLimiter {
	r := &RateLimiter{
		attempts: make(map[string]int),
		max:      max,
	}
	r.cond = sync.NewCond(&r.mu)
	return r
}

func (r *RateLimiter) Request(from string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for r.attempts[from] >= r.max {
		return fmt.Errorf("maximum request limit reached")
	}
	r.attempts[from]++
	go r.decrementAfterDelay(from)
	return nil
}

func (r *RateLimiter) decrementAfterDelay(from string) {
	time.Sleep(time.Second * 30)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.attempts[from]--
	r.cond.Broadcast()
}
