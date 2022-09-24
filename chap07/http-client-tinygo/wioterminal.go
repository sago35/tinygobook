package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/sago35/tinygo-examples/wioterminal/initialize"
	"tinygo.org/x/drivers/net/http"
)

var (
	ssid     string // 設定が必要
	password string // 設定が必要
)

func main() {
	//initialize.Debug(true)
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
	addr string // 設定が必要 (例: `192.168.1.102:8088`)
)

func run() error {
	res, err := http.Get(fmt.Sprintf("http://%s/", addr))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%s\r\n", body)

	return nil
}
