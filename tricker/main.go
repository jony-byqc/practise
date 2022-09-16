package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	var testChan = make(chan int, 2)
	//defer ticker.Stop()
	go sendToChan(testChan)
	for {
		select {
		case <-ticker.C:

		case <-testChan:
			fmt.Printf("123234")

		}
	}

	time.Sleep(10 * time.Second)
}
func sendToChan(testChan chan int) {
	for {
		time.Sleep(5 * time.Second)
		testChan <- 1
	}

}
