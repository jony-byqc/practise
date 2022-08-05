package main

import "fmt"

func main() {
	//leble:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break //leble
			}
			fmt.Println("j=", j)
		}
		fmt.Println("i=", i)
	}
	//fmt.Println("j=",j)
}
