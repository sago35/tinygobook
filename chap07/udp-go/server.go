package main

import (
	"fmt"
	"log"
	"net"
)

var (
	port int // 設定が必要 (例: 8088)
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listen: :%d\n", port)
	buf := [1024]byte{}
	for {
		n, _, _ := conn.ReadFromUDP(buf[:])
		fmt.Printf("from client: %q\r\n", string(buf[:n]))
	}
}
