package main

import (
	"flag"
	"io"
)

var (
	dataPathPtr = flag.String("d", "data",
		"data file path")

	queueNamePtr = flag.String("q", "queue",
		"queue name")
)

type DiskQueue struct {
	Queue  *Queue
	Buffer *Buffer
}

func NewDiskQueue() *DiskQueue {
	queue := &DiskQueue{
		Queue:  NewQueue(),
		Buffer: NewBuffer(),
	}
	return queue
}
func main() {
	d := NewDiskQueue()
	io.Copy(d.Queue, d.Buffer)
	_ = io.MultiWriter(d.Queue, d.Buffer)
	//io.Copy(mirror, d.Buffer)
}
