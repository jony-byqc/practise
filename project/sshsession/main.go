//package main
//
//import (
//	"fmt"
//	"golang.org/x/crypto/ssh"
//	"golang.org/x/crypto/ssh/terminal"
//	"log"
//	"net"
//	"os"
//	"time"
//)
//
//func main() {
//	// 可以使用 password 或者 sshkey 2种方式来认证。
//	sshHost := "192.168.60.157" // 主机名
//	sshUser := "wjq"            //用户名
//	sshPassword := "123456"     //密码
//	sshType := "password"       //ssh认证类型
//	//sshKeyPath := ""          //ssh id_rsa.id路径
//	sshPort := 22
//	//fmt.Print("请输入主机地址:")
//	//fmt.Scanln(&sshHost)
//	//fmt.Print("请输入主机端口:")
//	//fmt.Scanln(&sshPort)
//	//fmt.Print("请输入主机用户:")
//	//fmt.Scanln(&sshUser)
//	//fmt.Print("请输入主机密码:")
//	//fmt.Scanln(&sshPassword)
//
//	// 创建ssh登陆配置
//	config := &ssh.ClientConfig{
//		Timeout: time.Second, //ssh 连接timeout时间一秒钟，如果ssh验证错误 会在1秒内返回
//		User:    sshUser,     //指定ssh连接用户
//		// HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以，但是不够安全
//		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
//			return nil
//		},
//	}
//
//	if sshType == "password" {
//		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
//	}
//
//	// dial获取ssh Client
//	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
//	sshClient, err := ssh.Dial("tcp", addr, config)
//	if err != nil {
//		log.Fatal("创建ssh client 失败", err)
//	}
//	defer sshClient.Close()
//
//	// 创建ssh-session
//	session, err := sshClient.NewSession()
//	if err != nil {
//		log.Fatal("创建ssh session 失败", err)
//	}
//	defer session.Close()
//	// 将当前终端的stdin文件句柄设置给远程给远程终端，这样就可以使用tab键
//	fd := int(os.Stdin.Fd())
//	state, err := terminal.MakeRaw(fd)
//	if err != nil {
//		panic(err)
//	}
//	defer terminal.Restore(fd, state)
//
//	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
//	session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
//	session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备
//
//	// 设置终端模式
//	modes := ssh.TerminalModes{
//		ssh.ECHO:          0,     //禁止回显 （0 禁止,1 启动）
//		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
//		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
//	}
//
//	// 请求伪终端
//	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
//		log.Fatalf("request pty error: %s", err.Error())
//	}
//
//	// 启动远程shell
//	if err = session.Shell(); err != nil {
//		log.Fatalf("start shell error: %s", err.Error())
//	}
//
//	// 等待远程命令（终端）退出
//	if err = session.Wait(); err != nil {
//		log.Fatalf("return error: %s", err.Error())
//	}
//}

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"time"
)

func connect(user, password, host, key string, port int, cipherList []string) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	if key == "" {
		auth = append(auth, ssh.Password(password))
	} else {
		pemBytes, err := ioutil.ReadFile(key)
		if err != nil {
			return nil, err
		}

		var signer ssh.Signer
		if password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password))
		}
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}
