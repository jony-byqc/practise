package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		date = time.Now()
	)

	date = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
	date = date.Add(-10 * time.Minute)
	fmt.Println(date)
}
