package handlers

import (
	"encoding/json"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type OrdersHandler struct {
	serverHandler *ServerHandler
	db            *db.Cassandra
	pipeline      *server.Pipeline
}

func NewOrdersHandler(serverHandler *ServerHandler, db *db.Cassandra, pipeline *server.Pipeline) *OrdersHandler {
	return &OrdersHandler{
		serverHandler: serverHandler,
		db:            db,
		pipeline:      pipeline,
	}
}

func (handler *OrdersHandler) AddNewOrder(w http.ResponseWriter, req *http.Request) {
	handler.pipeline.Log <- server.PostNewOrder
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)

	if err != nil {
		handler.onError(w, http.StatusBadRequest, server.InvalidJSONDecoding, err)
		return
	}

	err = handler.db.AddOrder(order)
	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.OrderAddFailed+order.AwbNumber, err)
		return
	}

	message := server.OrderAdded + order.AwbNumber
	handler.pipeline.Log <- message
	io.WriteString(w, message)
}

func (handler *OrdersHandler) GetAllOrders(w http.ResponseWriter, req *http.Request) {
	handler.pipeline.Log <- server.GetOrders
	orders := handler.db.GetAllOrders()
	jsonData, err := json.Marshal(models.Orders{Orders: orders})
	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.InvalidJSONEncoding, err)
		return
	}
	length, err := w.Write(jsonData)
	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.DataSendFailed, err)
		return
	}
	handler.pipeline.Log <- strconv.Itoa(length)
}

func (handler *OrdersHandler) GetOrderById(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["id"]
	handler.pipeline.Log <- server.GetOrderWithId + id
	order, err := handler.db.GetOrdersById(id)
	if order.IsEmpty() || err != nil {
		handler.onError(w, http.StatusNotFound, server.OrderWithIdNotFound+id, err)
		return
	}
	jsonData, err := json.Marshal(order)
	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.InvalidJSONEncoding, err)
		return
	}
	length, err := w.Write(jsonData)
	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.DataSendFailed, err)
		return
	}
	handler.pipeline.Log <- server.SendOrderWithId + id + strconv.Itoa(length)
}

func (handler *OrdersHandler) onError(w http.ResponseWriter, status int, message string, err error) {
	e := models.ServerError{Status: status,
		ClientErrorMessage: message,
		Error:              err}
	handler.serverHandler.HandleError(w, e)
}
