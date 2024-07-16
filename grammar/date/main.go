package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取本月第一天的时间戳
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	firstDayTimestamp := firstDay.Unix()
	// 获取本月最后一天的时间戳
	lastDay := time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, time.Local)
	lastDayTimestamp := lastDay.Unix()

	fmt.Println(firstDayTimestamp, lastDayTimestamp)
}
