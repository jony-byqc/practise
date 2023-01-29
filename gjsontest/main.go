package main

import "fmt"

func main() {
	h, m, s := resolveTime(66666)
	fmt.Println(h, m, s)
}

// 秒转换为日时分秒和补零操作
func resolveTime(seconds int) (hour, minute, second int) {
	var day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return
}
