package main

import (
	"net"
	"log"
	"time"
	"sync"
	"fmt"
)

var clientArr = make(map[string]net.Conn)
var wg_server sync.WaitGroup

func main() {
	listen_socket, err := net.Listen("tcp4", "localhost:110")
	if err != nil {
		log.Println(err)
	}
	defer listen_socket.Close()

	wg_server.Add(3)

	//接受Client消息
	go func() {
		for {
			new_conn, err := listen_socket.Accept()
			clientArr[new_conn.RemoteAddr().String()] = new_conn
			if err != nil {
				continue
			}
			fmt.Println(new_conn.RemoteAddr().String() + " 上线了")
			go sendClient(new_conn)
		}
		wg_server.Done()
	}()

	//广播给所有Client发消息
	go func() {
		for {
			go clientListen()
			time.Sleep(time.Second * 10)
		}
		wg_server.Done()
	}()

	//给单个Client发消息
	go func() {
		for {
			var s string;
			fmt.Println("请输入客户端标识：")
			fmt.Scanf("%s", &s)
			go oneClientListen(s)
		}
		wg_server.Done()
	}()

	wg_server.Wait()

}

func sendClient(new_conn net.Conn) {
	buf := make([]byte, 500)

	for {
		n, err := new_conn.Read(buf)
		if err != nil {
			log.Println(err)
			//移除已经关闭的客户端，维护客户端队列
			removeClient(new_conn)
			new_conn.Close()
			return
		}
		log.Println(new_conn.RemoteAddr().String()+" clien msg:", string(buf[0:n]))
	}
}

func removeClient(new_conn net.Conn) {
	log.Println(clientArr)
	log.Println(new_conn.RemoteAddr().String() + " 已经阵亡")
	delete(clientArr, new_conn.RemoteAddr().String())
	log.Println("delete close conn")
	log.Println(clientArr)
	return
}

func clientListen() {
	if len(clientArr) > 0 {
		for _, v := range clientArr {
			v.Write([]byte("所有人休息会！！！！额要啪啪啪"))
		}
	} else {
		log.Println("no connection " + string(len(clientArr)))
	}
}

func oneClientListen(s string) {
	if _, ok := clientArr[s]; ok {
		//存在
		clientArr[s].Write([]byte("单独找你，papapapapappa"))
	} else {
		log.Println(s + " 已经***过度而去")
	}
}
