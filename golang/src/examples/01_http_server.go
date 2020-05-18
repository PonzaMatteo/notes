package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":12345"
	log.Println("start listening ", addr)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(addr, nil)
	panic(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %s %s", r.Method, r.URL.String())
	_, err := w.Write([]byte("HELLO!"))
	if err != nil {
		log.Print("failed to send response")
	}
}