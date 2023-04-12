package main

import (
	"bytes"
	dqueue "github.com/nsqio/go-diskqueue"
	"log"
	"sync"
	"time"
)

type Queue struct {
	dqueue.Interface
}

func logger(lvl dqueue.LogLevel, msg string, v ...any) {
	log.Printf(msg, v...)
}

func NewQueue() *Queue {

	return &Queue{
		Interface: dqueue.New(
			"queue", "data", 1<<20, 0, 65535, 1, time.Second, logger),
	}
}

func (b *Queue) Write(buf []byte) (int, error) {

	err := b.Interface.Put(buf)
	if err != nil {
		return 0, err
	}

	return len(buf), nil
}

type Buffer struct {
	bytes.Buffer
	mu sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{
		Buffer: bytes.Buffer{},
		mu:     sync.Mutex{},
	}
}

func (b *Buffer) Write(buf []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	w, err := b.Buffer.Write(buf)
	if err != nil {
		return 0, err
	}

	return w, nil
}

func (b *Buffer) Read(buf []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.Buffer.Read(buf)
}

func (b *Buffer) Len() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.Buffer.Len()
}
