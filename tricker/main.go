package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDDHHMISS = "2006-01-02 15:04:05"
)

type MyTicker struct {
	*time.Ticker               //扩展定时器
	interval     time.Duration //定时周期
	fn           func()        //回调
	chn          chan bool     //关闭信号
	status       bool          //状态      true表示启动 false表示非启动
}

//设置状态值
func (m *MyTicker) setStatus(status bool) {
	m.status = status

}

//获取状态值
func (m *MyTicker) getStatus() bool {
	return m.status
}

//NewTicker interval秒级周期 fn回调函数
func NewTicker(interval int64, fn func()) *MyTicker {
	m := &MyTicker{
		interval: time.Duration(interval) * time.Second,
		fn:       fn,
		chn:      make(chan bool),
	}
	return m
}

//Stop 关闭定时器
func (m *MyTicker) Stop() {
	fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "关闭定时器...") //打印
	if !m.getStatus() {
		fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "定时已经关闭") //打印
		return
	}

	//发送关闭信号
	m.chn <- true
}

//Stop 启动定时器
func (m *MyTicker) Start() {
	fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "启动定时器...") //打印
	if m.getStatus() {
		fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "定时已经开启") //打印
		return
	}

	//启动携程监听timerporc调度 tiker 以及 自定义关闭信号
	go func() {
		//启动 ticker
		m.Ticker = time.NewTicker(m.interval)
		m.setStatus(true)
		fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "定时启动") //打印

		//跳出for循环时 关闭ticker
		defer m.Ticker.Stop()
		defer m.setStatus(false)
		defer fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "定时关闭") //打印

		//阻塞监听调度信号
		for {
			select {
			case <-m.Ticker.C:
				//监听ticker 信号 调用任务
				go m.fn()
			case <-m.chn:
				//监听信号  跳出for 执行defer
				return
			default:

			}
		}
	}()
}

func main() {
	ticker := NewTicker(1, func() {
		fmt.Println(time.Now().Format(YYYYMMDDHHMISS), "定时器执行")
	})

	ticker.Stop() //检测 stop未启动ticker
	time.Sleep(10 * time.Second)
	ticker.Start() //检测 start
	time.Sleep(10 * time.Second)
	ticker.Start() //检测 start已启动ticker
	time.Sleep(10 * time.Second)
	ticker.Stop() //检测 stop
	time.Sleep(10 * time.Second)
	ticker.Start() //检测 重启开启ticker
	time.Sleep(10 * time.Second)
	ticker.Stop() //检测 重新关闭ticker
	time.Sleep(10 * time.Second)
	select {}
}

//defer导致 定时启动与定时关闭时间一致
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [select (no cases)]:// 说明go携程全部回收
//main.main()

func main1() {
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
