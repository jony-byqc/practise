package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1 // 大写字符串
	LOWER = 2 // 小写字符串
	CAP   = 4 // 字符串单词首字母大写
	REV   = 8 // 反转字符串
)

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样  //奇数为1，偶数为0
			array = append(array, i)
		}
	}
	return array
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}

func main1() {
	fmt.Printf("even: %v\n", even(100))
	a, b := swap(100, 200)
	fmt.Printf("swap: %d\t%d\n", a, b)
	fmt.Printf("shifting: %d\n", shifting(100))
	fmt.Printf("nagation: %d\n", nagation(100))
}
func main() {
	var a int = 3
	var b int = 3
	c := a & b
	fmt.Println(c)
	decimalToBinary(a)
	fmt.Println(procstr("HELLO PEOPLE!", LOWER|REV|CAP))
}

func procstr(str string, conf byte) string {
	// 反转字符串
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// 查询配置中的位操作
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}

func decimalToBinary(num int) {

	var binary []int

	for num != 0 {

		binary = append(binary, num%2)

		num = num / 2

	}

	if len(binary) == 0 {

		fmt.Printf("%d\n", 0)

	} else {

		for i := len(binary) - 1; i >= 0; i-- {

			fmt.Printf("%d", binary[i])

		}

		fmt.Println()

	}

}
