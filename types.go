package main

import (
	"log"
	"time"
)

type RateLimiter struct {
	C    chan struct{}
	stop chan struct{}

	ticker *time.Ticker
}

func NewFixedWindowRateLimiter(interval time.Duration, limit int) *RateLimiter {
	r := &RateLimiter{
		ticker: time.NewTicker(interval),
		C:      make(chan struct{}, limit),
		stop:   make(chan struct{}),
	}

	go func() {
		for {
			select {
			case tick := <-r.ticker.C:
				//todo len не проваряют
				for len(r.C) > 0 {
					<-r.C
					log.Println("clear", time.Since(tick))
				}
			case <-r.stop:
				return
			}
		}
	}()

	return r
}

func NewTokenBucketRateLimiter(interval time.Duration, limit int) *RateLimiter {
	refillRate := interval / time.Duration(limit)
	r := &RateLimiter{
		ticker: time.NewTicker(refillRate),
		C:      make(chan struct{}, limit),
		stop:   make(chan struct{}),
	}
	for i := 0; i < limit; i++ {
		r.C <- struct{}{}
	}

	go func() {
		for {
			select {
			case <-r.ticker.C:
				r.refill()
			case <-r.stop:
				return
			}
		}
	}()

	return r
}

// только для fixed window
func (r *RateLimiter) Consume() {
	r.C <- struct{}{}
	log.Println("consume", time.Now())
}

// только для token bucket
func (r *RateLimiter) Acquire() {
	<-r.C
	log.Println("acquire", time.Now())
}

func (r *RateLimiter) Stop() {
	r.ticker.Stop()
	close(r.stop)
	close(r.C)
}

func (r *RateLimiter) refill() {
	select {
	case r.C <- struct{}{}:
		log.Println("add")
	default:
	}
}
