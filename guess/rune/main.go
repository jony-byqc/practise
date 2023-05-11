package main

import (
	"fmt"
)

func main() {
	// 这样正常
	const r = 'a'

	var q int = r

	// 这样不能编译
	const d rune = 'a'

	var w int = r

	fmt.Println(q, w)

	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	//	m['a'] = 3
	fmt.Println(len(m))
}
