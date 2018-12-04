package main

import (
	"./db"
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
	Db 			   *db.Cassandra
}

func (server *Server) Start() {
	server.createBasicRoutes()
	if server.Port == "" {
		server.Port = DEFAULT_PORT
	}
	go server.run()
	server.connectDB()
}

func (server *Server) run() {
	fmt.Println("Server is running on port :", server.Port)
	var err = http.ListenAndServe(":" + server.Port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *Server) connectDB() {
	server.Db = &db.Cassandra{}
	server.Db.ConnectToCluster()
}

func (server *Server) createBasicRoutes() {
	http.HandleFunc("/", helloWorldHandler)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	var server = Server{Host: DEFAULT_HOST, Port: os.Getenv("PORT")}
	server.Start()
}
