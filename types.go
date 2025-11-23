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

// только для fixed window
func (r *RateLimiter) Consume() {
	r.C <- struct{}{}
	log.Println("consume", time.Now())
}

func (r *RateLimiter) Stop() {
	r.ticker.Stop()
	close(r.stop)
	close(r.C)
}
