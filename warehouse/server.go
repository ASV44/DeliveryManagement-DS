package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "8080"
	DEFAULT_TYPE = "http"
)

type Server struct {
	Host           string
	Port           string
	ConnectionType string
}

func (server *Server) Start() {
	server.createBasicRoutes()
	fmt.Println("Server is running on port :", server.Port)
	var err = http.ListenAndServe(server.Host+":"+server.Port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *Server) createBasicRoutes() {
	http.HandleFunc("/", helloWorldHandler)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	var server = Server{Host: DEFAULT_HOST, Port: DEFAULT_PORT}
	server.Start()
}
