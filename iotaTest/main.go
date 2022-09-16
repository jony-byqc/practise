package main

import (
	"fmt"
)

const (
	a = iota
	b1
	b2 = 5
	b3
	b4 = iota
	b5
	b6
	b7 = 1
	b8
	b9 = iota
	b10
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {

	fmt.Println(a, b1, b2, b3, b4, b5, b6, b7, b8, b9, b10)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
