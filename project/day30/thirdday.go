package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

const (
	a uint16 = 100
)

func main() {
	//s := NewService()
	//go func() {
	//	t := <-s.stopCh
	//	fmt.Println(t, "fsdaf")
	//}()

	//s.Run()
	//time.Sleep(1 * time.Hour)
	// 使用位掩码创建一个16位的二进制数
	//var aa uint16
	//
	//// 一次性设置多个位为1
	//for index := 0; index < int(6); index++ {
	//	aa = aa | (1 << index)
	//}
	//
	//fmt.Println(aa)
	//return

	//a := []int{1, 2, 3}
	//b := []int{4, 5, 6}
	//copy(a, b)
	//copy(b, a)
	//fmt.Println(a) // Output: [4 5 6]
	//fmt.Println(b) // Output: [1 2 3]

	//z := 9 >> 3 & 0b1
	//x := 3 & 0b1100111111111110
	//c := 9 & 0b1100111111111110
	//v := 15 & 0b1100111111111110
	//b := 31 & 0b1100111111111110
	//fmt.Println(z, x, c, v, b)

	//s := Shift(3, 4)
	//a := ShiftA(3, 4)
	//b := ShiftB(3, 4)
	//c := ShiftC(3, 4)
	//fmt.Println(s, a, b, c)

	//b := a | 0

	b := "" + string(1)
	c := []byte(b)
	fmt.Println(b, c)
}

type Service struct {
	stopCh chan struct{}
}

func NewService() *Service {
	return &Service{
		stopCh: make(chan struct{}),
	}
}

func (s *Service) Run() {

	select {
	case s.stopCh <- struct{}{}:
		fmt.Println("aaaaaaaaaaaaa")
	default:
		logrus.Error("sss")
	}
}

func Shift(value uint16, offset int) uint16 {
	switch value {
	case 0:
		return 1 << offset
	case 1:
		return 1 << (1 + offset)
	case 2:
		return 1 << (2 + offset)
	case 3:
		return 1 << (3 + offset)
	case 4:
		return 1 << (4 + offset)
	default:
		return value
	}
}

func ShiftA(value uint16, offset int) uint16 {
	return value << uint16(offset)
}

func ShiftB(value uint16, offset int) uint16 {
	return uint16(math.Pow(2, float64(offset))) * value
}

func ShiftC(value uint16, offset int) uint16 {
	return value << uint16(offset+int(value))
}
