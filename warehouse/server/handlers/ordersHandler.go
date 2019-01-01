package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server"
	"io"
	"net/http"
)

type OrdersHandler struct {
	db       *db.Cassandra
	pipeline *server.Pipeline
}

func NewOrdersHandler(db *db.Cassandra, pipeline *server.Pipeline) *OrdersHandler {
	return &OrdersHandler{
		db: db,
		pipeline: pipeline,
	}
}

func (handler *OrdersHandler) AddNewOrder(w http.ResponseWriter, req *http.Request) {
	handler.pipeline.Log <- "POST of new order"
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "invalid json")
		return
	}

	err = handler.db.AddOrder(order)
	var message = "Successfully added order : %s"
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Error while adding order : %s " + err.Error()
	}
	message = fmt.Sprintf(message, order.AwbNumber)
	handler.pipeline.Log <- message
	io.WriteString(w, message)
}