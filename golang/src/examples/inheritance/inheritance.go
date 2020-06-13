package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type handler struct {
	//embedding the http.ServeMux allow to use all its the methods
	//on a `handler` end extend it with new behaviour
	*http.ServeMux
}

func NewHandler() *handler {
	return &handler{http.DefaultServeMux}
}

//Method that allow to add as handler a function that return (value, error) and send the response\
//in json.
func (h *handler) HandleJson(patter string, serve func(r *http.Request) (interface{}, error)) {
	h.HandleFunc(patter, func(w http.ResponseWriter, r *http.Request) {
		resp, err := serve(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp = NewError(err)
		}
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
}

func main() {
	handler := NewHandler()
	//Note: HandleFunc is defined on *ServeMux
	handler.HandleFunc("/hello", hello)
	handler.HandleJson("/hello-json", helloJson)
	//Note: ListenAndServe accept an interface http.Handler that our handler inherit from the embedded *ServeMux
	err := http.ListenAndServe(":12345", handler)
	panic(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	message, err := GetMessage(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(NewError(err))
		return
	}
	_ = json.NewEncoder(w).Encode(message)
}

func helloJson(r *http.Request) (interface{}, error) {
	return GetMessage(r.RequestURI)
}

//Used for sending error as json
type Error struct {
	Error string
}
func NewError(err error) *Error {
	return &Error{Error: err.Error()}
}

//Simulate method that can fail
func GetMessage(str string) (map[string]string, error) {
	if rand.Float32() >= 0.5 {
		return nil, fmt.Errorf("failed to [...]")
	}
	return map[string]string{"msg": str}, nil
}
