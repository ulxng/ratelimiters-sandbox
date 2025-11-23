package main

import (
	"fmt"
	"time"
)

func main() {
	rl := NewFixedWindowRateLimiter(5*time.Second, 1)

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
