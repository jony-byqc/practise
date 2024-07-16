package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时区
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}

	// 获取当前时间
	now := time.Now().In(loc)

	// 获取 UTC 时间
	nowUTC := now.UTC()

	// 计算时间差
	timeDiff := now.Sub(nowUTC)

	fmt.Println("Current local time:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Current UTC time:", nowUTC.Format("2006-01-02 15:04:05"))
	fmt.Println("Time difference:", timeDiff)
}
