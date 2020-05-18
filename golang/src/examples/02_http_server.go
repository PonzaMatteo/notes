package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Food struct {
	Name string
	Icon string
}

var menu = []Food{
	{"Fries", "🍟"},
	{"Apple", "🍏"},
	{"Avocado", "🥑"},
	{"Pizza", "🍕"},
}

func main() {
	http.HandleFunc("/menu", Menu)
	err := http.ListenAndServe(":12345", nil)
	panic(err)
}

func Menu(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %s %s", r.Method, r.URL.String())
	err := json.NewEncoder(w).Encode(menu)
	if err != nil{
		log.Printf("failed to send response: %v", err)
	}
}
