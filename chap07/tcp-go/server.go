package main

import (
	"fmt"
	"log"
	"net"
)

var (
	serverIP string // 設定が必要 (例: `192.168.1.102`)
	port     int    // 設定が必要 (例: 8088)
)

func main() {
	conn, err := net.Listen(`tcp`, fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("listen: %s:%d\n", serverIP, port)
	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Printf("from client: %q\r\n", string(buf[:n])) // 受信したメッセージを返す
	conn.Write(buf[:n])
	conn.Close()
}
