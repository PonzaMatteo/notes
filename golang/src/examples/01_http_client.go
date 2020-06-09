package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	path := "hello"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	resp, err := http.Get("http://localhost:12345/" + path)
	if err != nil {
		log.Printf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	fmt.Print(string(content))
}
