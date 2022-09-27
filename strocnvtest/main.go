package main

import (
	"fmt"
	"strconv"
)

func main() {
	ID, _ := strconv.Atoi("0001")
	a := fmt.Sprintf("%08b\n", ID) //001000
	fmt.Println(a)
	string2Bin()
	fmt.Println(converseToBin(16))
	main2()
}
func string2Bin() {
	i, err := strconv.ParseInt("100000011", 2, 64)
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

func main2() {
	s := "123456789" //英文和数字和英文特殊字符占1个字节长度,中文占用3个字节长度
	a := s[2]
	fmt.Println(a)        //输出:
	fmt.Printf("%T\n", a) //输出uint8
	b := fmt.Sprintf("%c", a)
	fmt.Printf("%T\n", b) //输出:string
	fmt.Println(b)        //输出
	fmt.Println(len(s))   //输出:
	fmt.Println(s[1:4])   //输出:
	fmt.Println(s[:2])    //输出:
	fmt.Println(s[5:8])   //输出:
	fmt.Println(s[2:5])   //输出:
}
