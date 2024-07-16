package main

import (
	"fmt"
	"time"
)

type stu struct {
	start chan struct{}
	stop  chan struct{}
}

func newStu() *stu {
	return &stu{
		stop:  make(chan struct{}),
		start: make(chan struct{}, 1),
	}
}

func main() {

	s := newStu()
	s.start <- struct{}{}
	go s.run()
	s.stop <- struct{}{}
	time.Sleep(2 * time.Second)

	fmt.Println("333333333333")
}

func (s *stu) run() {
	select {
	case <-s.stop:
		fmt.Println("11111111111")
	case <-s.start:
		fmt.Println("222222222222")

	}
}
