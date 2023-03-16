package main

import (
	"flag"
	"fmt"
)

func main() {

	//定义变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" ,就是 -u 指定参数
	//"" , 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名,默认为空") //'\"asdf\"'保留双引号
	flag.StringVar(&pwd, "p", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3300, "端口号，默认为3300")
	//转换
	flag.Parse()
	//输出结果
	//	fmt.Printf("+++++++++++++++++++ %v\n", user)
	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user, pwd, host, port)

}
