package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(str string)
}

type FileReader struct {
	Folder
}

func (f *FileReader) Read() string {
	return "file content"
}

func (f *FileReader) Write(str string) {
	fmt.Println("not supported")
}

type BookReader struct {
}

type FlipEr interface {
	Flip() string
}

type Folder interface {
	Fold(str string)
}

func (b *BookReader) Flip() string {
	return "file content"
}

func (b *BookReader) Fold(str string) {
	fmt.Println("not supported")
}

func main() {
	var r Reader = &FileReader{}
	w, ok := r.(Writer)
	if ok {
		reader, ok1 := w.(Folder)
		if ok1 {
			fmt.Println(reader)
			fmt.Println("Book接口类型转换成功")
		} else {
			fmt.Println("Book接口类型转换失败")
		}
	} else {
		fmt.Println("接口类型转换失败")
	}
}
