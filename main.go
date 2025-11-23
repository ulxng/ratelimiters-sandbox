package main

import (
	"fmt"
	"time"
)

func main() {
	rl := NewTokenBucketRateLimiter(5*time.Second, 3)

	start := time.Now()
	rl.Acquire()
	rl.Acquire()
	rl.Acquire()
	rl.Acquire()
	rl.Acquire()
	rl.Acquire()
	rl.Acquire()

	fmt.Println(time.Since(start))
	rl.Stop()
}
