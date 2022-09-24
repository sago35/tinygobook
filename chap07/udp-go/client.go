package main

import (
	"fmt"
	"net"
	"time"
)

var (
	port int // 設定が必要 (例: 8088)
)

func main() {
	conn, _ := net.Dial("udp", fmt.Sprintf("255.255.255.255:%d", port))
	for i := 0; i < 3; i++ {
		fmt.Printf("send to 255.255.255.255:%d\n", port)
		fmt.Fprintf(conn, "udp from go %d\r\n", i)
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}
