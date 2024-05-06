package main

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"time"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`

	Addresss string `json:"addresss"`
}

func main() {
	f1()
	t := time.Now().UnixNano()
	m := make(map[string]interface{})
	person := Person{
		Name:    "zhangsan",
		Address: "北京海淀",
	}

	elem := reflect.ValueOf(&person).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		name := relType.Field(i).Name
		if elem.Field(i).String() != "" {
			m[name] = elem.Field(i).String()
		}
	}

	fmt.Println(m)
	fmt.Println(fmt.Sprintf("反射-duration:%d", time.Now().UnixNano()-t))
}

type UserInfo struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Agee string `json:"agee"`
}

func f1() {
	u1 := UserInfo{Name: "q1mi", Age: "18"}
	m3 := structs.Map(&u1)
	for k, v := range m3 {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
	fmt.Println(m3)

}
