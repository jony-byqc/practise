package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 [min, max) 范围内的随机浮点数
	min := 0.0
	max := 10.0
	randFloat := rand.Float64()*(max-min) + min

	fmt.Println(randFloat)
}
