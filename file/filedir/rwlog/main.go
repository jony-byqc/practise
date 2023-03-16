package main

import (
	"bufio"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/hpcloud/tail/ratelimiter"
	"gopkg.in/tomb.v1"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Tail struct {
	Filename string
	Lines    chan *Line
	Config

	tomb.Tomb // provides: Done, Kill, Dying
	// contains filtered or unexported fields
}

type Config struct {
	// File-specifc
	//	Location    *SeekInfo // Seek to this location before tailing
	ReOpen      bool // Reopen recreated files (tail -F)
	MustExist   bool // Fail early if the file does not exist
	Poll        bool // Poll for file changes instead of using inotify
	Pipe        bool // Is a named pipe (mkfifo)
	RateLimiter *ratelimiter.LeakyBucket

	// Generic IO
	Follow      bool // Continue looking for new lines (tail -f)
	MaxLineSize int  // If non-zero, split longer lines into multiple lines

	// Logger, when nil, is set to tail.DefaultLogger
	// To disable logging: set field to tail.DiscardingLogger
	//Logger logger
}

type Line struct {
	Text string
	Time time.Time
	Err  error // Error from tail
}

func main() {
	//main1()
	fileName := `./log/a.log`
	config := tail.Config{
		ReOpen:    true,                                 // true则文件被删掉阻塞等待新建该文件，false则文件被删掉时程序结束
		Follow:    true,                                 // true则一直阻塞并监听指定文件，false则一次读完就结束程序
		Location:  &tail.SeekInfo{Offset: 0, Whence: 0}, // Location读取文件的位置, Whence更加系统选择参数从哪开始：0从头，1当前，2末尾
		MustExist: true,                                 // 允许日志文件不存在
		Poll:      true,                                 // 轮询
	}
	// 打开文件读取日志
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	// 开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)
	if tails.Lines == nil {
		log.Println("tails.Lines 为空，return")
		return
	}
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			//time.Sleep(time.Second) // 读取出错停止一秒
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}

func main1() {
	file, err := os.Open("./log/a.log")
	if err != nil {
		log.Fatalf("Open file fail:%v", err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		lineStr := strings.TrimSpace(string(line))
		l3 := strings.Count(lineStr, "") - 1 // 计算""在lineStr中非重叠个数
		if l3 < 5 {
			file.Seek(0, os.SEEK_END)
		}
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}
		//超过50个字符的数据行会被处理
		if len(string(line)) >= 50 {
			//处理代码
		}
	}
}
