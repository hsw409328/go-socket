package main

import (
	"net"
	"log"
	"fmt"
	"sync"
	"time"
	"os"
)

var wg_client sync.WaitGroup

func main() {
	addr := "localhost:110"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		log.Print(err)
		log.Print("建立TcpAddr对象失败")
		//退出程序
		os.Exit(1)
	}
	connTcp, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Print(err)
		log.Print("建立Server通迅失败")
		os.Exit(1)
	}

	wg_client.Add(2)

	//向Server发送消息
	go func() {
		for {
			var s string;
			fmt.Scanf("%s", &s)
			connTcp.Write([]byte(s))
		}
		wg_client.Done()
	}()

	//监听Server通知
	go func() {
		for {
			go listenServerMessage(connTcp)
			time.Sleep(time.Second * 10)
		}
		wg_client.Done()
	}()

	wg_client.Wait()
}

func listenServerMessage(connTcp *net.TCPConn) {
	buf := make([]byte, 1024)
	n, err := connTcp.Read(buf)
	if err != nil {
		log.Print(err)
		connTcp.Close()
		os.Exit(1)
		return
	}
	log.Print("server message: ", string(buf[0:n]))
}
