package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 定义日期格式的正则表达式
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}$`)

	// 要检查的字符串
	dateString := "1999-03"

	// 使用正则表达式检查字符串是否符合格式
	if dateRegex.MatchString(dateString) {
		fmt.Println("字符串符合2024-03格式")
	} else {
		fmt.Println("字符串不符合2024-03格式")
	}
}
