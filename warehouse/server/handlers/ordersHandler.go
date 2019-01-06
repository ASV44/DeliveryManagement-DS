package handlers

import (
	"encoding/json"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/mappers"
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

	if order.IsEmpty() || err != nil {
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

func (handler *OrdersHandler) RegisterNewOrders(w http.ResponseWriter, req *http.Request) {
	var orders models.Orders
	err := json.NewDecoder(req.Body).Decode(&orders)

	if orders.Orders == nil || err != nil {
		handler.onError(w, http.StatusBadRequest, server.InvalidJSONDecoding, err)
		return
	}
	handler.pipeline.Log <- server.PostNewOrders + strconv.Itoa(len(orders.Orders))

	insertErrors := handler.db.RegisterNewOrders(orders.Orders)
	if insertErrors != nil {
		handler.serverHandler.HandleInsertErrors(w, insertErrors)
		return
	}

	message := server.RegisterMultipleOrders
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
	order, err := handler.db.GetOrderById(id)
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
	handler.pipeline.Log <- server.SendOrderWithId + id
	handler.pipeline.Log <- strconv.Itoa(length)
}

func (handler *OrdersHandler) GetOrdersByAWB(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	awbNumber := variables["awb_number"]
	handler.pipeline.Log <- server.GetSpecificAWBOrders + awbNumber
	orders := handler.db.GetOrdersByAWB(awbNumber)
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
	handler.pipeline.Log <- server.SendSpecificAWBOrders + awbNumber
	handler.pipeline.Log <- strconv.Itoa(length)
}

func (handler *OrdersHandler) DeleteOrder(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersHandler) UpdateOrder(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["id"]
	handler.pipeline.Log <- server.UpdateOrderRequest + id
	values := mappers.UrlQueryToMap(req.URL.Query(), mappers.FieldToOrderColumn)
	err := handler.db.UpdateOrderById(id, values)

	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.OrderUpdateFailed+id, err)
		return
	}

	order, _ := handler.db.GetOrderById(id)
	jsonData, _ := json.Marshal(order)
	length, err := w.Write(jsonData)

	if err != nil {
		handler.onError(w, http.StatusInternalServerError, server.DataSendFailed, err)
		return
	}
	handler.pipeline.Log <- server.UpdatedOrderWithId + id
	handler.pipeline.Log <- strconv.Itoa(length)
}

func (handler *OrdersHandler) onError(w http.ResponseWriter, status int, message string, err error) {
	e := models.ServerError{Status: status,
		ClientErrorMessage: message,
		Error:              err.Error()}
	handler.serverHandler.HandleError(w, e)
}
