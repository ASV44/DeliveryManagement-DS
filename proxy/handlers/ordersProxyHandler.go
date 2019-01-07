package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/go-redis/redis"
	"net/http"
)

type OrdersProxyHandler struct {
	pipeline *proxy.Pipeline
}

func NewOrdersProxyHandler(pipeline *proxy.Pipeline, cache *redis.Client) *OrdersProxyHandler {
	return &OrdersProxyHandler{pipeline: pipeline}
}

func (handler *OrdersProxyHandler) AddNewOrder(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) RegisterNewOrders(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) GetAllOrders(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) GetOrderById(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) GetOrdersByAWB(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) UpdateOrder(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) DeleteOrder(w http.ResponseWriter, req *http.Request) {

}

func (handler *OrdersProxyHandler) DeleteMultipleOrder(w http.ResponseWriter, req *http.Request) {

}
