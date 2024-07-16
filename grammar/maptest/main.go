package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[string]int)

func main() {
	// 创建一个普通的 map

	// 在单线程中对 map 进行操作
	myMap["apple"] = 5

	// 读取 map 中的值
	fmt.Println(myMap["apple"])  // 输出: 5
	fmt.Println(myMap["banana"]) // 输出: 3
	fmt.Println(myMap["orange"]) // 输出: 7

	// 删除 map 中的键值对
	//delete(myMap, "banana")

	go z()
	go d()

	// 遍历 map
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}
	// 输出:
	// apple: 5
	// orange: 7
	var myMaps sync.Map
	myMaps.Store("key1", "value1")
	myMaps.Store("key2", "value2")
	value, ok := myMaps.Load("key1")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}

func z() {
	myMap["banana"] = 3
	myMap["orange"] = 7
}

func d() {
	delete(myMap, "banana")
}
