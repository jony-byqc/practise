package main

import "fmt"

func main() {
	x := 10
PositiveNumber:
	fmt.Println("x is positive")

BigNumber:
	fmt.Println("x is greater than 5")
	if x > 0 {
		goto PositiveNumber
	}

	if x > 5 {
		goto BigNumber
	}

	fmt.Println("x is between 0 and 5")
	return

}
