package server

import (
	"../db"
	"../models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
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
	db       *db.Cassandra
	pipeline chan string
}

func New(host string, port string) server {
	if host == "" {
		host = DefaultHost
	}
	if port == "" {
		port = DefaultPort
	}

	return server{
		host: host,
		port: port,
	}
}

func (server *server) Start() {
	server.pipeline = make(chan string)
	server.createBasicRoutes()
	go server.run()
	server.connectDB()
	server.processPipeline()
}

func (server *server) run() {
	fmt.Println("server is running on port :", server.port)
	var err = http.ListenAndServe(":"+server.port, server.router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *server) connectDB() {
	server.db = &db.Cassandra{}
	server.db.ConnectToCluster()
}

func (server *server) processPipeline() {
	for {
		select {
		case data := <-server.pipeline:
			fmt.Println(data)
		}
	}
}

func (server *server) createBasicRoutes() {
	server.router = mux.NewRouter()
	server.router.HandleFunc("/", server.helloWorldHandler)
	server.router.HandleFunc("/order", server.addNewOrder).Methods("POST")
}

func (server *server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	server.pipeline <- "request on '/' route"
	io.WriteString(w, "Hello world!")
}

func (server *server) addNewOrder(w http.ResponseWriter, req *http.Request) {
	server.pipeline <- "POST of new order"
	decoder := json.NewDecoder(req.Body)
	var order models.Order
	err := decoder.Decode(&order)
	if err != nil {
		panic(err)
	}

	err = server.db.AddOrder(order)
	var message = "Successfully added order : %s"
	if err != nil {
		message = "Error while adding order : %s " + err.Error()
	}
	message = fmt.Sprintf(message, order.AwbNumber)
	fmt.Println(message)
	io.WriteString(w, message)
}
