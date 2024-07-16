package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	add()

	for {
		select {}

	}
}

func add() {
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(logrus.New())))
	_, err := c.AddFunc("0 23 31 * *", func() {
		fmt.Println("111111")
	})
	if err != nil {
	}
	c.Start()
}
