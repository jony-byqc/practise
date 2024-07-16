package main

import (
	"fmt"
	"strings"
)

// 打印列表a中没有列表b中的字段
func printMissingInA(a, b []string) {
	bMap := make(map[string]bool)

	// 将列表b中的字符串放入map中
	for _, v := range b {
		bMap[v] = true
	}

	// 遍历列表a并检查是否在map中存在
	for _, v := range a {
		if _, found := bMap[v]; !found {
			fmt.Println("列表a中缺少的字段:", v)
		}
	}
}

// 打印列表b中没有列表a中的字段
func printMissingInB(a, b []string) {
	aMap := make(map[string]bool)

	// 将列表a中的字符串放入map中
	for _, v := range a {
		aMap[v] = true
	}

	// 遍历列表b并检查是否在map中存在
	for _, v := range b {
		if _, found := aMap[v]; !found {
			fmt.Println("列表b中缺少的字段:", v)
		}
	}
}

func main() {
	str := "apple,,banana,cherry,,date"
	splitStr := strings.Split(str, ",")
	fmt.Println(splitStr) // 输出: [apple  banana cherry  date]

	listA := []string{"apple", "banana", "orange", "grape"}
	listB := []string{"banana", "orange", "watermelon", "kiwi"}

	fmt.Println("列表a中没有列表b中的字段:")
	printMissingInA(listA, listB)

	fmt.Println("\n列表b中没有列表a中的字段:")
	printMissingInB(listA, listB)
}
