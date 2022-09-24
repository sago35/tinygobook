package main

import (
	"fmt"
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
	cnt = 0
)

func handler(w http.ResponseWriter, req *http.Request) {
	cnt++
	fmt.Fprintf(w, "hello %d", cnt)
	fmt.Printf("%d %#v\r\n", cnt, req.UserAgent())
}

func run() error {
	fmt.Printf("listen: %s:80\r\n", initialize.IP().String())
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":80", nil)
}
