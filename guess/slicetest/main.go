package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//1不见了，导致1不见的原因是当调用append时，将创建一个新切片。新切片具有新的 “长度” 属性，该属性不是指针，但Data属性仍指向同一个底层数组。因此，我们函数内的代码最终会反转切片所引用的底层数组（切片里边是不存储任何数据的），但是函数外原始切片的长度属性还是之前的长度值3，这就是造成了上面 1 被丢掉的原因。
func reverse1(s []int) {

	s = append(s, 999)

	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//当我们调用append时，会创建一个新的切片。在第二个例子中，反转函数里的新切片仍指向同一底层数组，因为数组有足够的容量来添加新元素，因此在函数内对底层数组的更改也能在函数外体现，但是这个例子中，在reverse函数里向切片添加了三个元素，而此时我们的切片的底层数组没有足够的容量来添加新元素了，于是系统分配了一个新数组，让切片指向该数组。这时函数内外的切片指向的不同的底层数组，所以在函数内对切片做的任何更改都不会再影响我们的初始切片。
func reverse2(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
