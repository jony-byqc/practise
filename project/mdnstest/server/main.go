package main

import (
	"github.com/grandcat/zeroconf"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//"_workstation._tcp"定义了服务名称
	//参数形式：
	//func Register(instance, service, domain string, port int, text []string, ifaces []net.Interface)

	server, err := zeroconf.Register("GoZeroconf", "a_workstation._tcp", "local.", 42424, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
		// Exit by user
	case <-time.After(time.Second * 120):
		// Exit by timeout
	}

	log.Println("Shutting down.")
}
