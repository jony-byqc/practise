package main

import (
	"fmt"
	"os"
)

func main() {
	//fileObej, err := os.Open("./log/a.log") // 读取文件，可以使用相对路径
	ReadFile("./log/a.log")
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
