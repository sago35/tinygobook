package main

import (
	"io"
	"log"
	"os"
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

	for i := 0; i < 10; i++ {
		err = run()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}

	select {}
}

func run() error {
	res, err := http.Get("http://tinygo.org")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
	println("\r")
	return nil
}
