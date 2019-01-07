package router

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy/handlers"
	"github.com/gorilla/mux"
)

func NewProxyRouter(serverHandler *handlers.ProxyServerHandler,
	ordersHandler *handlers.OrdersProxyHandler) *mux.Router {
	router := mux.NewRouter()
	addBasicRoutesHandlers(router, serverHandler)
	addOrdersRoutesHandlers(router, ordersHandler)

	return router
}

func addBasicRoutesHandlers(router *mux.Router, handler *handlers.ProxyServerHandler) {
	router.HandleFunc("/", handler.ProxyRootHandler)
}

func addOrdersRoutesHandlers(router *mux.Router, handler *handlers.OrdersProxyHandler) {
	router.HandleFunc("/order", handler.AddNewOrder).Methods("POST")
	router.HandleFunc("/order/{id}", handler.GetOrderById).Methods("GET")
	router.HandleFunc("/order/{id}", handler.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/{id}", handler.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/orders", handler.RegisterNewOrders).Methods("POST")
	router.HandleFunc("/orders", handler.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders/{awb_number}", handler.GetOrdersByAWB).Methods("GET")
	router.HandleFunc("/orders", handler.DeleteMultipleOrder).Methods("DELETE")
}
