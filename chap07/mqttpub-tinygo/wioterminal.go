package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/sago35/tinygo-examples/wioterminal/initialize"
	"tinygo.org/x/drivers/net/mqtt"
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
	broker    = "tcp://test.mosquitto.org:1883"
	topic     = "tinygobook"
	clientfmt = "tinygo-client-%X"
)

func run() error {
	// ClientIDのためのランダム値を生成する
	var rndbuf [4]byte
	rand.Read(rndbuf[:])

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(fmt.Sprintf(clientfmt, rndbuf))

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	defer c.Disconnect(250)

	for i := 1; ; i++ {
		// MQTTにPublish(送信)する
		text := fmt.Sprintf("mqtt from tinygo %X %d", rndbuf, i)
		fmt.Printf("send : %s\r\n", text)
		token := c.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
	return nil
}
