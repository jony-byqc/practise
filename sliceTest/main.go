package main

import "fmt"

func main1() {
	var s []int
	fmt.Println(len(s))
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse1(s)
	fmt.Println(s)
	var a []int
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	reverse(a)
	fmt.Println(a)
}
func reverse1(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var a [][]float32

	fmt.Println(a)
}
