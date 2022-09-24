package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	addr string // 設定が必要 (例: `192.168.1.102:8088`)
)

func main() {
	res, err := http.Get(fmt.Sprintf("http://%s/", addr))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
