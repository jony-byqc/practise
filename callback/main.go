package main

//import (
//	"fmt"
//	"strconv"
//)
//
////一个记录日志的类型：func(string)
//type saveLog func(msg string)
//
////将字符串转换为int64,如果转换失败调用saveLog
//func stringToInt(s string, log saveLog) int64 {
//
//	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
//		log(err.Error())
//		return 0
//	} else {
//		return value
//	}
//}
//
////记录日志消息的具体实现
//func myLog(msg string) {
//	fmt.Println("Find Error:", msg)
//}
//
//func miniourl() {
//	fmt.Println(stringToInt("123", myLog)) //转换时将调用mylog记录日志
//	fmt.Println(stringToInt("s", myLog))
//}

import "fmt"

type Callback func(x, y int) int

//提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
	return callback(x, y)
}

func add(x, y int) int {
	return x + y
}

func main() {
	x, y := 6, 7
	z := test(x, y, add)
	fmt.Printf("z = %d", z)
}
