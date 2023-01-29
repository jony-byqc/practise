package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd() //获取当前目录
	if err != nil {
		fmt.Println("get dir error:", err)
	}
	fmt.Println(dir)
	fmt.Println("开始创建一个文本文件test.json")
	if f, err := os.Create("D:/photo/src/jsonfile.json"); err == nil {
		fmt.Println("文件创建成功。。。")
		// 将要写入文本的内容
		_, err = f.Write([]byte("{123}"))
		if err == nil {
			fmt.Println("文件写入内容成功。。。")
		}
	} else {
		fmt.Println("文件创建失败！")
	}
}
