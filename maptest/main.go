package main

import "fmt"

func main() {
	// map 判断key值是否存在 判断方式为value,key := map[key], key为true则存在
	demo := map[string]string{
		"name":  "tom",
		"phone": "010-xxxx",
	}

	demo2 := map[string]bool{
		"hot": true,
		"top": false,
	}

	fmt.Println("demo1[\"name\"]: ", demo["name"]) //tom
	fmt.Println("demo1[\"sex\"]: ", demo["sex"])   //无输出

	fmt.Println("demo2[\"hot\"]: ", demo2["hot"]) //true
	fmt.Println("demo2[\"top\"]: ", demo2["top"]) //false  判断方式错误 top存在 但是返回值为false

	_, name := demo["name"]
	fmt.Println("is exist demo1[\"name\"] ?", name) //true

	_, sex := demo["sex"]
	fmt.Println("is exist demo1[\"sex\"] ?", sex) //false

	if _, hot := demo2["hot"]; hot {
		fmt.Println("is exist demo2[\"hot\"] ?", hot) //true
	}

	if _, top := demo2["hot"]; top {
		fmt.Println("is exist demo2[\"top\"] ?", top) //true
	}

	if _, old := demo2["old"]; old {
		fmt.Println("is exist demo2[\"old\"] ?", old)
	} else {
		fmt.Println("is exist demo2[\"old\"] ?", old) //false
	}

	intMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	stringMap := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	boolMap := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}

	interfaceMap := map[string]interface{}{
		"a": true,
		"b": "b",
		"c": 1,
	}

	x := interfaceMap["w"]
	fmt.Println(x)
	fmt.Println(intMap[1], intMap[5])
	fmt.Println(stringMap["1"], stringMap["5"])
	fmt.Println(len(stringMap["1"]), len(stringMap["5"])) // len=1,len=0
	fmt.Println(boolMap["a"], boolMap["c"], boolMap["e"])
	fmt.Println(interfaceMap["a"], interfaceMap["c"], interfaceMap["e"])

	if value, ok := stringMap["6"]; ok {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在, value为空值:", value)
	}

}
