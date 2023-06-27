package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"unsafe"
)

//func main() {
//	// 这里假设你的乱码文字是以 GBK 编码存储的
//	gbkData := []byte("��������:5Nm")
//
//	decoded, err := DecodeGBK(gbkData)
//	if err != nil {
//		fmt.Println("解码失败:", err)
//		return
//	}
//
//	fmt.Println(decoded) // 输出：链家中国
//}
//
//func DecodeGBK(data []byte) (string, error) {
//	decoder := simplifiedchinese.GBK.NewDecoder()
//	decoded, err := decoder.Bytes(data)
//	if err != nil {
//		return "", err
//	}
//
//	return string(decoded), nil
//}

//func main() {
//	s := "��������" // 假设这个是乱码的文字
//	decoded, _ := decode([]byte(s))
//	fmt.Println(decoded) // 输出：中文世界
//}
//
//func decode(data []byte) (string, error) {
//	// 如果这个乱码文字的编码是 UTF-8，直接转换为字符串
//	if utf8.Valid(data) {
//		return string(data), nil
//	}
//
//	// 如果是其他编码，你需要确定该乱码文字的实际编码并进行解码
//	// 例如，如果是 GBK 编码，可以使用 golang.org/x/text/encoding 包进行解码
//
//	return "", fmt.Errorf("Unknown encoding")
//}

//import (
//	"fmt"
//	"strings"
//)
//import "github.com/axgle/mahonia"
//
//func main() {
//	enc := mahonia.NewEncoder("gbk")
//	//converts a  string from UTF-8 to gbk encoding.
//	fmt.Println(enc.ConvertString("��������"))
//	s := coverString("��������")
//	fmt.Println(s)
//}
//
//func coverString(src string) string {
//	return replaceNullHtml(coverGBKToUTF8(src))
//}
//
//func replaceNullHtml(src string) string {
//	temp := strings.Replace(src, "聽", "", -1)
//	return temp
//}
//
//func coverGBKToUTF8(src string) string {
//	return mahonia.NewDecoder("gbk").ConvertString(src)
//}

func main() {
	d := bytes.Buffer{}
	b := []byte{179, 204, 208, 242, 195, 251, 179, 198}
	//bs := []byte{179, 204, 208, 242, 195, 251, 179, 198}
	d.Write(b)
	fmt.Println(d.String())

	//fmt.Println(bs)
	//st := "测试拧紧:5Nm"
	//stb := []byte(st)
	//fmt.Println(stb)
	//s := Bytes2String(b)
	//by := String2Bytes(s)
	//fmt.Println(by)

	//a := []byte{69, 0, 0, 30, 118, 209, 0, 0, 128, 17, 0, 0, 192, 168, 31, 49, 39, 97, 233, 183}
	//fmt.Println(fmt.Sprintf("%x", a))
	//c := HexToBytes(fmt.Sprintf("%x", b))
	//fmt.Println(string(c))
	//b := []byte{69, 0, 0, 30, 118, 209, 0, 0, 128, 17, 0, 0, 192, 168, 31, 49, 39, 97, 233, 183}
	//fmt.Println(fmt.Sprintf("%x", b))
	//c := HexToBytes(fmt.Sprintf("%x", b))
	//fmt.Println(string(c))
}

func HexToBytes(hexStr string) []byte {
	hexBytes, err := hex.DecodeString(hexStr)
	if err == nil {
		return hexBytes
	} else {
		return nil
	}
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
