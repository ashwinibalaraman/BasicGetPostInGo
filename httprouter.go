package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	mux.POST("/hello", helloPost)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

type User struct {
	Name   string
	Gender string
	Age    int
	Id     int
}

type ReqParameters struct {
	Name string
}

type ResParameters struct {
	Greeting string
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	ReqParams := ReqParameters{}

	ReqParams.Name = "ash"
	json.NewDecoder(req.Body).Decode(&ReqParams)

	ResParams := ResParameters{
		Greeting: "Hello, " + ReqParams.Name + "!",
	}

	greeting, _ := json.Marshal(ResParams)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s\n", greeting)
}
