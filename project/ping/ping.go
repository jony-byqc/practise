package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"sync"
)

func ping(ip string, pong chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("ping", ip)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ping %s 失败, err:%s\n", ip, err)
		pong <- fmt.Sprintf("%s  ====>  %v", ip, false)
		return
	}
	output, err := simplifiedchinese.GBK.NewDecoder().Bytes(out)
	if err != nil {
		fmt.Printf("编码 %+v 失败, err:%s\n", out, err)
		pong <- fmt.Sprintf("%s  ====>  %v", ip, false)
		return
	}
	info := string(output)
	fmt.Printf("ping %s 成功, 返回信息:%s\n", ip, info)
	exist := strings.Contains(info, "TTL=")
	if !exist {
		pong <- fmt.Sprintf("%s  ====>  %v", ip, false)
		return
	}
	pong <- fmt.Sprintf("%s  ====>  %v", ip, true)
}

func main() {
	prex := "192.168.10."
	start := 10
	end := 254
	result := make(chan string, end-start)
	wg := sync.WaitGroup{}
	wg.Add(end - start)
	for i := start; i < end; i++ {
		go ping(fmt.Sprintf("%s%d", prex, i), result, &wg)
	}
	wg.Wait()
	close(result)
	for v := range result {
		fmt.Printf("%s\n", v)
	}
}
