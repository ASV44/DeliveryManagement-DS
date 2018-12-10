package main

import (
	"./db"
	"./models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "8080"
)

type Server struct {
	host           string
	port           string
	router		   *mux.Router
	db 			   *db.Cassandra
	pipeline	   chan string
}

func main() {
	var server = Server{host: DEFAULT_HOST, port: os.Getenv("PORT")}
	server.Start()
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
	var err = http.ListenAndServe(":" + server.port, server.router)
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
	server.router = mux.NewRouter()
	server.router.HandleFunc("/", server.helloWorldHandler)
	server.router.HandleFunc("/order", server.addNewOrder).Methods("POST")
}

func (server *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	server.pipeline <- "request on '/' route"
	io.WriteString(w, "Hello world!")
}

func (server *Server) addNewOrder(w http.ResponseWriter, req *http.Request) {
	server.pipeline <- "POST of new order"
	decoder := json.NewDecoder(req.Body)
	var order models.Order
	err := decoder.Decode(&order)
	if err != nil {
		panic(err)
	}

	err = server.db.AddOrder(order)
	var message = "Successfully added order : "
	if err != nil {
		message = "Error while adding order : "
	}
	fmt.Println(message + order.AwbNumber, err)
	io.WriteString(w, message + order.AwbNumber + "\n" + err.Error())
}
