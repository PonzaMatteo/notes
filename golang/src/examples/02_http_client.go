package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Food struct {
	Name string
	Icon string
}

func main() {
	start := time.Now()
	defer func() {
		log.Println("total running time: ", time.Now().Sub(start).Milliseconds(), "ms")
	}()
	resp, err := http.Get("http://localhost:12345/menu")
	if err != nil {
		log.Printf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close() // ignoring error
	var food []Food
	err = json.NewDecoder(resp.Body).Decode(&food)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	fmt.Println(food)
}
