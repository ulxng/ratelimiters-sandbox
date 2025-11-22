package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	C    chan struct{}
	stop chan struct{}

	ticker *time.Ticker
}

func NewRateLimiter(interval time.Duration, limit int) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(interval),
		C:      make(chan struct{}, limit),
		stop:   make(chan struct{}),
	}
}

func (r *RateLimiter) FixedWindow() {
	go func() {
		for {
			select {
			case tick := <-r.ticker.C:
				for len(r.C) > 0 {
					<-r.C
					fmt.Println("clear", time.Since(tick))
				}
			case <-r.stop:
				r.ticker.Stop()
				close(r.C)
				return
			}
		}
	}()
}

func (r *RateLimiter) Consume() {
	r.C <- struct{}{}
	fmt.Println("consume", time.Now())
}

func (r *RateLimiter) Stop() {
	close(r.stop)
}
