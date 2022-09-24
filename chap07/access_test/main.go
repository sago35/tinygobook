package main

import (
	"fmt"
	"time"

	"github.com/sago35/tinygo-examples/wioterminal/initialize"
)

var (
	ssid     string // 設定が必要
	password string // 設定が必要
)

func main() {
	//initialize.Debug(true)
	rtl, err := initialize.SetupRTL8720DN()
	if err != nil {
		failMessage(err.Error())
	}

	err = rtl.ConnectToAccessPoint(ssid, password, 10*time.Second)
	if err != nil {
		failMessage(err.Error())
	}

	ip, subnet, gateway, err := rtl.GetIP()
	if err != nil {
		failMessage(err.Error())
	}
	for {
		fmt.Printf("IP Address : %s\r\n", ip)
		fmt.Printf("Mask       : %s\r\n", subnet)
		fmt.Printf("Gateway    : %s\r\n", gateway)
		time.Sleep(10 * time.Second)
	}
}

func failMessage(str string) {
	for {
		fmt.Printf("%s\n", str)
		time.Sleep(5 * time.Second)
	}
}
