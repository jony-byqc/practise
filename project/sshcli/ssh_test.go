package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSshSimple(t *testing.T) {
	username := "wjq"
	password := "123456"
	ip := "192.168.60.157"
	port := "22"
	client := NewSSHClient(username, password, ip, port)
	// 1.运行远程命令

	err := client.RunTerminal("top", os.Stdout, os.Stdin)
	if err != nil {
		return
	}

	cmd := "ls"
	backinfo, err := client.Run(cmd)
	if err != nil {
		fmt.Printf("failed to run shell,err=[%v]\n", err)
		return
	}
	fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// 2. 上传一文件
	filename := "main.go"
	client.UploadFile(filename, filename)
	// 上传
	n, err := client.UploadFile(filename, "/tmp/"+filename)
	if err != nil {
		fmt.Printf("upload failed: %v\n", err)
		return
	}
	// 3. 显示该文件
	cmd = "cat " + "/tmp/" + filename
	backinfo, err = client.Run(cmd)
	if err != nil {
		fmt.Printf("run cmd faild: %v\n", err)
		return
	}
	fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// 4. 下载该文件到本地
	n, err = client.DownloadFile("/tmp/"+filename, filename+"download")
	if err != nil {
		fmt.Printf("download failed: %v\n", err)
		return
	}
	fmt.Printf("download file[%v] ok, size=[%d]\n", filename, n)
}
