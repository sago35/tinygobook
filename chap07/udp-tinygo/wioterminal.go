package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sago35/tinygo-examples/wioterminal/initialize"
	"tinygo.org/x/drivers/net"
)

var (
	ssid     string // 設定が必要
	password string // 設定が必要
)

func main() { //initialize.Debug(true)
	_, err := initialize.Wifi(ssid, password, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = run()
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

var (
	port int // 設定が必要 (例: 8088)
)

func run() error {
	ip := net.ParseIP("255.255.255.255")
	raddr := &net.UDPAddr{IP: ip, Port: port}
	laddr := &net.UDPAddr{Port: port}

	conn, err := net.DialUDP("udp", laddr, raddr)
	for ; err != nil; conn, err = net.DialUDP("udp", laddr, raddr) {
		time.Sleep(5 * time.Second)
	}

	// 3回送信する
	for i := 0; i < 3; i++ {
		fmt.Printf("send to 255.255.255.255:%d\r\n", port)
		fmt.Fprintf(conn, "udp from tinygo %d\r\n", i)
		time.Sleep(1 * time.Second)
	}

	// 受信を待ち受ける
	fmt.Printf("listen : %d\r\n", port)
	buf := [1024]byte{}
	for {
		n := int(0)
		for {
			n, err = conn.Read(buf[:])
			if err != nil {
				return err
			}
			if n > 0 {
				break
			}
			time.Sleep(5 * time.Millisecond)
		}
		fmt.Printf("recv: %q\r\n", string(buf[:n]))
	}

	conn.Close()
	return nil
}
