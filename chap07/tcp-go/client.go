package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	serverIP string // 設定が必要 (例: `192.168.1.102`)
	port     int    // 設定が必要 (例: 8088)
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}
	// サーバーに送信する
	fmt.Fprintf(conn, "hello from go\r\n")
	// サーバーからの受信を標準出力に書き出す
	io.Copy(os.Stdout, conn)
	conn.Close()
}
