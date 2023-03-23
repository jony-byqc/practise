package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Skill string
}

func ms() {
	stu := Student{"tom", 12, "football"}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return
	}
	fmt.Println("序列化后: ", string(data))
}
func map2json() {
	// map 转 Json字符串
	m := make(map[string]interface{})
	m["name"] = "jetty"
	m["age"] = 16

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return
	}
	fmt.Println("序列化后: ", string(data)) // 打印: 序列化后:  {"age":16,"name":"jetty"}

	// Json字符串 转 map
	str := `{"age":25,"name":"car"}`
	err = json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
		return
	}
	fmt.Printf("反序列化后: map=%v, name=%v\n", m, m["name"])
	// 打印: 反序列化后: map=map[age:25 name:car], name=car
	//如果我们传入一个结构体指针，那么 JSON 序列化会使用指针所指向的结构体的值作为序列化的源头，也就是说，如果原始的结构体值在序列化之后被改变了，那么 JSON 序列化的结果也会发生变化。
	//
	//其次，当我们传入一个结构体实例时，JSON 序列化会将该结构体实例的值进行复制并序列化，不会影响原始结构体实例的值。这样做的好处是，即使原始结构体实例的值发生了改变，序列化的结果也不会受到影响。
	//
	//因此，如果我们希望在 JSON 序列化过程中保持原始结构体实例的值不变，请传入该结构体的实例。如果希望序列化的结果与原始结构体实例的值保持同步，请传入该结构体的指针。
}
