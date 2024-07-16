package main

import "fmt"

func main() {
	// 定义一个 ASCII 码数组
	asciiCodes := []int{1, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33, 32, 49, 50, 51, 52, 53, 54, 55, 56, 57, 33, 64, 35, 36, 37, 94, 38, 42, 40, 41, 95, 43}

	// 将 ASCII 码转换为字符串
	asciiStr := ""
	var a = make([]byte, 0)
	for _, code := range asciiCodes {
		asciiStr += string(rune(code))
		a = append(a, []byte(asciiStr)...)
	}

	fmt.Println("ASCII string:", asciiStr, a)
}
