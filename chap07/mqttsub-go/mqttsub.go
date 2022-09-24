package main

import (
	"crypto/rand"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	broker    = "tcp://test.mosquitto.org:1883"
	topic     = "tinygobook"
	clientfmt = "tinygobook-sub-%X"
)

func subHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("%s : %s\n", msg.Topic(), msg.Payload())
}

func main() {
	// ClientIDのためのランダム値を生成する
	var rndbuf [4]byte
	rand.Read(rndbuf[:])

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(fmt.Sprintf(clientfmt, rndbuf))

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Subscribe時の処理はgoroutineで実行されるためselectで待つ
	token := c.Subscribe(topic, 0, subHandler)
	token.Wait()
	select {}
}
