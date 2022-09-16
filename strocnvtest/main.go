package main

import (
	"fmt"
	"strconv"
)

func main() {
	ID, _ := strconv.Atoi("00000001")
	fmt.Printf("%08d\n", ID) //001000
	string2Bin()
	fmt.Println(converseToBin(000000010))
}
func string2Bin() {
	i, err := strconv.ParseInt("00000011", 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

func converseToBin(n int) string { //int2bin
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
