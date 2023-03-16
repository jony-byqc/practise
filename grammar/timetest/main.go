package main

import (
	"fmt"
	"time"
)

const audivwDatetimeLayout = "20/11/2011 11:34:45"

func main() {
	t := time.Now()
	fmt.Println(t)
	stime, err := time.Parse(audivwDatetimeLayout, t.Format(time.RFC3339))
	if err != nil { //nolint
		return
	}
	fmt.Println(stime)
}
