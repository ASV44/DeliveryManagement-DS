package proxy

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "8000"
)

type server struct {
	host     string
	port     string
	router   *mux.Router
	pipeline *Pipeline
}

func New(host string, port string, pipeline *Pipeline, router *mux.Router) server {
	if host == "" {
		host = DefaultHost
	}
	if port == "" {
		port = DefaultPort
	}

	return server{
		host:     host,
		port:     port,
		router:   router,
		pipeline: pipeline,
	}
}

func (server *server) Start() {
	go server.run(server.router)
	server.processPipeline()
}

func (server *server) run(router *mux.Router) {
	server.pipeline.Log <- "Proxy is running on port : " + server.port
	var err = http.ListenAndServe(":"+server.port, router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *server) processPipeline() {
	for {
		select {
		case log := <-server.pipeline.Log:
			fmt.Println(log)
		}
	}
}
