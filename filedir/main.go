package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//fileObej, err := os.Open("./log/a.log") // 读取文件，可以使用相对路径
	ReadFile("./log/a.log")
	src := "./YAML/c.yaml" //项目路径下要提前拥有该yaml文件
	writeToXml(src)
	readFromXml(src)
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Study struct {
	CourseName string `yaml:"CourseName"`
	Score      int    `yaml:"Score"`
}
type Student struct {
	Name      string  `yaml:"name"`
	Address   string  `yaml:"addr"`
	ScoreList []Study `yaml:"ScoreList"`
}

func writeToXml(src string) {
	stu := &Student{
		Name:      "George",
		Address:   "北京",
		ScoreList: []Study{{"语文", 21}, {"数学", 22}},
	}
	data, err := yaml.Marshal(stu) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	checkError(err)
	err = ioutil.WriteFile(src, data, 0777)
	checkError(err)
}
func readFromXml(src string) {
	content, err := ioutil.ReadFile(src)
	checkError(err)
	newStu := &Student{}
	err = yaml.Unmarshal(content, &newStu)
	checkError(err)
	ScoreList := newStu.ScoreList
	fmt.Println(newStu.Name + "的学习情况")
	for _, v := range ScoreList {
		fmt.Println("Course:" + v.CourseName + "\tScore:" + strconv.Itoa(v.Score))
	}
}

type Conf struct {
	Test []string `yaml:"array.test,flow"`
}

func main1() {
	data := `array.test: ["val1", "val2", "val3"]`
	var conf Conf
	yaml.Unmarshal([]byte(data), &conf)

	data2, _ := yaml.Marshal(conf)
	fmt.Printf("%s\n", string(data2))
}
