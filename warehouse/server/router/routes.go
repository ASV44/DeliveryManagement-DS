package router

import (
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server/handlers"
	"github.com/gorilla/mux"
)

func New(serverHandler *handlers.ServerHandler, ordersHandler *handlers.OrdersHandler) *mux.Router {
	router := mux.NewRouter()
	addBasicRoutesHandlers(router, serverHandler)
	addOrdersRoutesHandlers(router, ordersHandler)

	return router
}

func addBasicRoutesHandlers(router *mux.Router, handler *handlers.ServerHandler) {
	router.HandleFunc("/", handler.RootHandler)
}

func addOrdersRoutesHandlers(router *mux.Router, handler *handlers.OrdersHandler) {
	router.HandleFunc("/order", handler.AddNewOrder).Methods("POST")
	router.HandleFunc("/order/{id}", handler.GetOrderById).Methods("GET")
	router.HandleFunc("/orders", handler.RegisterNewOrders).Methods("POST")
	router.HandleFunc("/orders", handler.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders/{awb_number}", handler.GetOrdersByAWB).Methods("GET")
}
