package main

import (
	"fmt"
	"unsafe"
)

func MyFunction(i int, arr [2]int) {
	i = 22
	arr[1] = 99
	fmt.Printf("in my_function - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

type MyStruct struct {
	i int
}

func MyFunction2(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function2 - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

func BaseTypeTest() {
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	MyFunction(i, arr)
	fmt.Printf("after calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

func CompoundTypeTest() {
	a := MyStruct{i: 30}
	b := &MyStruct{i: 40}
	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	MyFunction2(a, b)
	fmt.Printf("after calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

type MyStruct2 struct {
	i, j int
}

func MyFunction3(ms *MyStruct2) {
	fmt.Printf("pointer adress in function: %p\n", &ms)
	ptr := unsafe.Pointer(ms)

	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8*i)))
		*c = *c + i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

// MyFunction3与3_1功能一样，3使用的直接修改内存地址的方式修改的指针指向结构体的字段值
func MyFunction3_1(ms *MyStruct2) {
	fmt.Printf("pointer adress in function: %p\n", &ms)
	ms.i = 41
	ms.j = 52

}

func MyFunction4(ms *MyStruct2) {
	fmt.Printf("pointer adress in function: %p\n", &ms)
	ms = &MyStruct2{
		i: 41,
		j: 52,
	}

	fmt.Printf("in my_function4 - [%p] %v\n", ms, ms)
}

//func MyFunction4 (ms *MyStruct2) {
//	ms
//}

func PointedDataUpdateTest() {
	a := &MyStruct2{i: 40, j: 50}
	fmt.Printf("[%p] %v  pointer adress outside: %p\n", a, a, &a)
	MyFunction3(a)
	fmt.Printf("[%p] %v\n", a, a)
}

func UpdatePointerDirection() {
	a := &MyStruct2{i: 40, j: 50}
	fmt.Printf("before calling my_function4 - [%p] %v\n", a, a)
	MyFunction4(a)
	fmt.Printf("after calling my_function4 - [%p] %v\n", a, a)
}

func main() {
	BaseTypeTest()
	fmt.Println("=======================分割线===================")
	CompoundTypeTest()
	fmt.Println("=======================分割线===================")
	PointedDataUpdateTest()
	fmt.Println("=======================分割线===================")
	UpdatePointerDirection()
}
