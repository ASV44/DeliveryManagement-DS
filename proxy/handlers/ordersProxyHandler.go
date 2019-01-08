package handlers

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

type OrdersProxyHandler struct {
	proxyRequestHandler *ProxyRequestHandler
	pipeline            *proxy.Pipeline
	cache               *redis.Client
}

func NewOrdersProxyHandler(proxyRequestHandler *ProxyRequestHandler,
	pipeline *proxy.Pipeline,
	cache *redis.Client) *OrdersProxyHandler {
	return &OrdersProxyHandler{proxyRequestHandler: proxyRequestHandler,
		pipeline: pipeline,
		cache:    cache}
}

func (handler *OrdersProxyHandler) AddNewOrder(w http.ResponseWriter, req *http.Request) {
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, proxy.ForwardOrderPost)
}

func (handler *OrdersProxyHandler) RegisterNewOrders(w http.ResponseWriter, req *http.Request) {
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, proxy.ForwardOrdersRegister)
}

func (handler *OrdersProxyHandler) GetAllOrders(w http.ResponseWriter, req *http.Request) {
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, proxy.ForwardOrdersGet)
}

func (handler *OrdersProxyHandler) GetOrderById(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["id"]
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, fmt.Sprintf(proxy.ForwardOrderGetById, id))
}

func (handler *OrdersProxyHandler) GetOrdersByAWB(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["awb_number"]
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, fmt.Sprintf(proxy.ForwardOrdersGetByAwb, id))
}

func (handler *OrdersProxyHandler) UpdateOrder(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["id"]
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, fmt.Sprintf(proxy.ForwardOrderUpdate, id))
}

func (handler *OrdersProxyHandler) DeleteOrder(w http.ResponseWriter, req *http.Request) {
	variables := mux.Vars(req)
	id := variables["id"]
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, fmt.Sprintf(proxy.ForwardOrderDelete, id))
}

func (handler *OrdersProxyHandler) DeleteMultipleOrder(w http.ResponseWriter, req *http.Request) {
	handler.proxyRequestHandler.forwardWarehouseRequest(w, req, proxy.ForwardOrdersDelete)
}
