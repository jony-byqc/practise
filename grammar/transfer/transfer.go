package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func DoubleHeight(rect Rectangle) {
	rect.Height = rect.Height * 2
}

func main() {
	rect := Rectangle{
		Width:  10,
		Height: 3,
	}

	// 这实际上不会修改 rect
	DoubleHeight(rect)

	fmt.Println("Width", rect.Width)
	fmt.Println("Height", rect.Height)
}
