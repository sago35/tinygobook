package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	addr string // 設定が必要 (例: `192.168.1.102:8088`)
	cnt  = 0
)

func handler(w http.ResponseWriter, req *http.Request) {
	cnt++
	fmt.Fprintf(w, "hello %d", cnt)
	fmt.Printf("%d %#v\n", cnt, req.UserAgent())
}

func main() {
	fmt.Printf("listen: %s\n", addr)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
