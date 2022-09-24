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
	serverIP string // 設定が必要 (例: `192.168.1.102`)
	port     int    // 設定が必要 (例: 8088)
)

func run() error {
	fmt.Printf("connect to %s:%d\r\n", serverIP, port)
	ip := net.ParseIP(serverIP)
	raddr := &net.TCPAddr{IP: ip, Port: port}
	laddr := &net.TCPAddr{Port: port}

	conn, err := net.DialTCP("tcp", laddr, raddr)
	for ; err != nil; conn, err = net.DialTCP("tcp", laddr, raddr) {
		time.Sleep(5 * time.Second)
	}

	if _, err = conn.Write([]byte("hello from tinygo\r\n")); err != nil {
		return err
	}

	buf := [1024]byte{}
	n := int(0)
	for {
		n, err = conn.Read(buf[:])
		if err != nil {
			return err
		}
		if n > 0 {
			break
		}
	}
	fmt.Printf("from server: %q\r\n", string(buf[:n]))

	fmt.Printf("Disconnecting TCP\r\n")
	conn.Close()

	return nil
}
