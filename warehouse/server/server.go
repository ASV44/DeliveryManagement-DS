package server

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
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
	db       *db.Cassandra
	pipeline chan string
}

func New(host string, port string, db *db.Cassandra) server {
	if host == "" {
		host = DefaultHost
	}
	if port == "" {
		port = DefaultPort
	}

	return server{
		host:     host,
		port:     port,
		db:       db,
		pipeline: make(chan string),
	}
}

func (server *server) Start() {
	router := server.createBasicRoutes()
	go server.run(router)
	server.processPipeline()
}

func (server *server) run(router *mux.Router) {
	fmt.Println("warehouse is running on port :", server.port)
	var err = http.ListenAndServe(":"+server.port, router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (server *server) processPipeline() {
	for {
		select {
		case data := <-server.pipeline:
			fmt.Println(data)
		}
	}
}

func (server *server) createBasicRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", server.helloWorldHandler)
	router.HandleFunc("/order", server.addNewOrder).Methods("POST")

	return router
}

func (server *server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	server.pipeline <- "request on '/' route"
	io.WriteString(w, "Hello world!")
}

func (server *server) addNewOrder(w http.ResponseWriter, req *http.Request) {
	server.pipeline <- "POST of new order"
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "invalid json")
		return
	}

	err = server.db.AddOrder(order)
	var message = "Successfully added order : %s"
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Error while adding order : %s " + err.Error()
	}
	message = fmt.Sprintf(message, order.AwbNumber)
	server.pipeline <- message
	io.WriteString(w, message)
}
