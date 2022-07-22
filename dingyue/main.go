package main

import (
	"fmt"
	_ "net/http/pprof"
	"strconv"
	"syscall"
	"unsafe"
)

func main() {
	var start int
	var end int
	fmt.Println("爬取酷我音乐盒mv ： https://www.kuwo.cn/mvplay/")
	fmt.Println("请输入开始id：")
	fmt.Scanln(&start)
	fmt.Println("请输入结束id：")
	fmt.Scanln(&end)
	SetCmdTitle("id:" + strconv.Itoa(start) + "-" + strconv.Itoa(end))

	go ProducerFunc(start, end) // kafka生产者
	go ConsumerFunc()           // kafka消费者

	select {} // 停止在cmd，避免cmd自动关闭
	// Init()
}

//func Init() {
//	router.RouterInit() // 启动http服务且初始化路由
//}

// 修改 cmd 的标题
func SetCmdTitle(title string) {
	kernel32, _ := syscall.LoadLibrary(`kernel32.dll`)

	sct, _ := syscall.GetProcAddress(kernel32, `SetConsoleTitleW`)

	syscall.Syscall(sct, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)

	syscall.FreeLibrary(kernel32)

}
