package main

import (
	"log"
	"testing"
)

const sudoPwPrompt = "sudo_password"

func TestSshSimple(t *testing.T) {
	username := "wjq"
	password := "123456"
	ip := "192.168.60.157"
	port := "22"
	client := NewSSHClient(username, password, ip, port)
	client.Connect()
	// 1.运行远程命令

	//err := client.RunTerminal("top", os.Stdout, os.Stdin)
	//if err != nil {
	//	return
	//}
	cmd := "cd /data"
	//cmd = "sudo -p " + sudoPwPrompt + " -S " + cmd
	//backinfo, err := client.Run(cmd)
	//if err != nil {
	//	fmt.Printf("failed to run shell,err=[%v]\n", err)
	//	return
	//}

	// 创建新的 SSH 会话
	session, err := client.sshClient.NewSession()
	// 创建一个缓冲区来存储远程命令的输出
	//var stdoutBuf bytes.Buffer
	//session.Stdout = &stdoutBuf
	if err != nil {
		log.Fatalf("创建 SSH 会话失败：%s", err)
	}
	defer session.Close()
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("session.StdinPipe() : %s", err)
		return
	}
	//_, err = session.StdinPipe()
	//if err != nil {
	//	log.Fatalf("创建 SSH 会话失败：%s", err)
	//}
	//modes := ssh.TerminalModes{
	//	ssh.ECHO:          1,     // disable echo
	//	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	//	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	//}

	// 指定远程工作目录
	//remoteDir := "/data"
	//o, err := session.CombinedOutput(fmt.Sprintf("cd %s", remoteDir))
	//if err != nil {
	//	log.Fatalf("切换远程目录失败：%s", err)
	//}

	//// 请求伪终端
	//if err = session.RequestPty("linux", 32, 160, modes); err != nil {
	//	log.Fatalf("request pty error: %s", err.Error())
	//}

	//session.Stdout = nil
	//session.Stderr = nil
	// 执行远程命令
	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("执行远程命令失败：%s", err)
	}

	write, err := stdinPipe.Write([]byte("ls"))
	if err != nil {
		return
	}

	log.Println(session.Stdout, write)

	//fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// 2. 上传一文件
	//filename := "main.go"
	//client.UploadFile(filename, filename)
	//// 上传
	//n, err := client.UploadFile(filename, "/tmp/"+filename)
	//if err != nil {
	//	fmt.Printf("upload failed: %v\n", err)
	//	return
	//}
	//// 3. 显示该文件
	//cmd = "cat " + "/tmp/" + filename
	//backinfo, err = client.Run(cmd)
	//if err != nil {
	//	fmt.Printf("run cmd faild: %v\n", err)
	//	return
	//}
	//fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	//// 4. 下载该文件到本地
	//n, err = client.DownloadFile("/tmp/"+filename, filename+"download")
	//if err != nil {
	//	fmt.Printf("download failed: %v\n", err)
	//	return
	//}
	//fmt.Printf("download file[%v] ok, size=[%d]\n", filename, n)
}
