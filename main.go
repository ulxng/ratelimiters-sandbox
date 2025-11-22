package main

import (
	"fmt"
	"time"
)

func main() {
	rl := NewRateLimiter(5*time.Second, 1)
	rl.FixedWindow()

	start := time.Now()
	rl.Consume()
	rl.Consume()
	rl.Consume()
	rl.Consume()
	rl.Consume()
	rl.Consume()
	rl.Consume()

	fmt.Println(time.Since(start))
	rl.Stop()
}
