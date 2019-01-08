package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/go-redis/redis"
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
	handler.proxyRequestHandler.forwardPostRequest(w, req, proxy.ForwardOrderPost)
}

func (handler *OrdersProxyHandler) RegisterNewOrders(w http.ResponseWriter, req *http.Request) {
	handler.proxyRequestHandler.forwardPostRequest(w, req, proxy.ForwardOrdersRegister)
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
