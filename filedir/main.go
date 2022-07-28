package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//fileObej, err := os.Open("./log/a.log") // 读取文件，可以使用相对路径
	ReadFile("./log/a.log")
	main1()
}
func ReadFile(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0400)
	var read_buffer = make([]byte, 10)
	var content_buffer = make([]byte, 0)
	fileinfo, _ := file.Stat()
	size := fileinfo.Size() //文件大小，单位是字节，int64
	var length int64 = 0    //标记已经读取了多少字节的内容
	for length < size {     //循环读取文件内容
		n, _ := file.Read(read_buffer)
		content_buffer = append(content_buffer, read_buffer[:n]...)
		length += int64(n)
	}
	fmt.Println(string(content_buffer))
}
func main1() {
	//pwd,_ := os.Getwd() //获取当前目录
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir("./log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].ModTime()) //打印当前文件或目录下的文件或目录名

	}
	filepath.Walk("./log", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)        //打印path信息
		fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}
