package main

import (
	"fmt"
	"time"
)

type S struct {
	S []float32 `json:"s"`
}
type T struct {
	T []float32 `json:"t"`
}

type W struct {
	W []float32 `json:"t"`
}
type M struct {
	M []float32 `json:"t"`
}

func main() {
	var a = []float32{}
	a = make([]float32, 1000)
	//b := float32(0)
	//for i := 1; i < 10; i++ {
	//	a = append(a, b)
	//}
	ti := time.Now()
	fmt.Println(ti)
	fmt.Println(a)
}
