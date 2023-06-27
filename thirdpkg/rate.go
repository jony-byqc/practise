package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// 每秒钟最多允许4次请求
	limiter := rate.NewLimiter(4, 10)

	for i := 0; i < 1000; i++ {
		if !limiter.Allow() {
			fmt.Printf("%d: throttle\n", i)
		} else {
			fmt.Printf("%d: do something\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}
