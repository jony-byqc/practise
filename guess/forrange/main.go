package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		s = append(s, v)
		fmt.Printf("len(s)=%v\n", len(s))
	}
}
