package main

import (
	"fmt"
	"strconv"
)

func main() {
	ID, _ := strconv.Atoi("001000")
	fmt.Printf("%06d\n", ID) //001000
}
