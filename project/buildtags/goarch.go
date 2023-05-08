package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS
	fmt.Println(runtime.GOARCH)
	fmt.Println(version)
	if sysType == "linux" {
		// LINUX系统
		fmt.Println("Linux system")
	}

	if sysType == "windows" {
		// windows系统
		fmt.Println("Windows system")
	}
}
