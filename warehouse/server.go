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
	host           string
	port           string
	connectionType string
	db 			   *db.Cassandra
	pipeline	   chan string
}

func (server *Server) Start() {
	server.pipeline = make(chan string)
	server.createBasicRoutes()
	if server.port == "" {
		server.port = DEFAULT_PORT
	}
	go server.run()
	server.connectDB()
	server.processPipeline()
}

func (server *Server) run() {
	fmt.Println("Server is running on port :", server.port)
	var err = http.ListenAndServe(":" + server.port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *Server) connectDB() {
	server.db = &db.Cassandra{}
	server.db.ConnectToCluster()
}

func (server *Server) processPipeline() {
	for {
		select {
		case data := <- server.pipeline:
			fmt.Println(data)
		}
	}
}

func (server *Server) createBasicRoutes() {
	http.HandleFunc("/", server.helloWorldHandler)
}

func (server *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	server.pipeline <- "request on / route"
	io.WriteString(w, "Hello world!")
}

func main() {
	var server = Server{host: DEFAULT_HOST, port: os.Getenv("PORT")}
	server.Start()
}
