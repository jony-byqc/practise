package main

import (
	"fmt"
	"time"
)

// WriteData
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据=%v\n", i)
		time.Sleep(500)
	}
	close(intChan) //关闭
}

//readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//读完数据后人物完成
	exitChan <- true
	close(exitChan)

}

func main() {

	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
