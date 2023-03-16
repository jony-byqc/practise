package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
)

func main() {

	byteValue := "aaaaaaaaaaaaaaaaaaaaaaaaaa"

	t := byteValue[:4]
	t1, _ := hex.DecodeString(t)
	torque := new(big.Int).SetBytes(t1)
	//IntValue2 := new(big.Int).SetInt64(1000000000000000000)
	fmt.Println("IntValue==>", torque)
	//fmt.Println("IntValue2==>", IntValue2)
	//fmt.Println(IntValue.Mul(IntValue, IntValue2))

	s := byteValue[4:8]
	s1, _ := hex.DecodeString(s)
	speed := new(big.Int).SetBytes(s1)
	fmt.Println("IntValue==>", speed)

	parseInt, err := strconv.ParseInt(byteValue[12:20], 16, 0)
	if err != nil {
		return
	}

	p, _ := strconv.ParseInt(ReverseString(strconv.FormatInt(parseInt, 2)), 2, 64)
	fmt.Println("IntValue==>", p)

}

func ReverseString(raw string) string {
	rt := ""
	for _, v := range raw {
		rt = string(v) + rt
	}
	return rt
}

var Website int32

func New() binary.ByteOrder {
	return binary.LittleEndian
}
func readRows() {
	file, err := os.Open("")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Data as hex: %x\n", data)
	//fmt.Printf("Data as string: %s\n", data)
	//fmt.Println("Number of bytes read:", len(data))
	m := Website
	var nember int
	for i := 1; i <= 100000; i++ {
		if i == 1 {
			nember = 100
		} else {
			nember = 14
		}
		data, err1 := readNextBytes(file, nember)
		if errors.Is(err1, io.EOF) {
			return
		}
		if len(data) == 100 {
			continue
		}
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		//readFull(file)
		fmt.Println("第", i, "个值为：", m)
		hexHeader := hex.EncodeToString(data)
		fmt.Println(hexHeader)

		if i > 100 {
			return
		}
	}
}

func readNextBytes(file *os.File, number int) ([]byte, error) {
	if number == 100 {
		bytes := make([]byte, 100)
		return bytes, nil
	}
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
		return nil, err
	}
	//f := convertDataArrayFromBytes(bytes, New())
	//fmt.Println(f)
	return bytes, nil
}

func convertDataArrayFromBytes(data []byte, order binary.ByteOrder) []float64 {

	length := len(data)
	//x := length % 4
	//if x != 0 {
	//	return ret
	//}

	ret := make([]float64, length)
	idx := 0
	for i := 0; i < length; i++ {
		reader := bytes.NewReader(data[0:])
		if err := binary.Read(reader, order, &ret[idx]); err != nil {
			fmt.Println("Error ")
		}
		idx += 1
	}
	return ret
}

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func readFull() {
	file, err := os.Open("D:\\github\\atcive\\practise\\file\\readbin\\20221012-121158-264.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	buff := make([]byte, 100)
	for {
		lens, err := file.Read(buff[10:])
		if err == io.EOF || lens < 0 {
			break
		}
	}
	hexHeader := hex.EncodeToString(buff)
	fmt.Println(hexHeader)
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
