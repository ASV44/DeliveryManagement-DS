package server

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "8080"
)

type server struct {
	host     string
	port     string
	router   *mux.Router
	pipeline *Pipeline
	db       *db.Cassandra
}

func New(host string, port string, pipeline *Pipeline, router *mux.Router, db *db.Cassandra) server {
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
		db:       db,
		pipeline: pipeline,
	}
}

func (server *server) Start() {
	go server.run(server.router)
	server.processPipeline()
}

func (server *server) run(router *mux.Router) {
	server.pipeline.Log <- "warehouse is running on port : " + server.port
	var err = http.ListenAndServe(":"+server.port, router)
	if err != nil {
		server.pipeline.Log <- err.Error()
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
