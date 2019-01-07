package main

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/ASV44/DeliveryManagement-DS/proxy/handlers"
	"github.com/ASV44/DeliveryManagement-DS/proxy/router"
	"os"
)

func main() {
	pipeline := proxy.InitPipeline()
	serverHandler := handlers.NewProxyServerHandler(pipeline)
	ordersProxyHandler := handlers.NewOrdersProxyHandler(pipeline)
	proxyServerRouter := router.NewProxyRouter(serverHandler, ordersProxyHandler)
	instance := proxy.New("", os.Getenv("PORT"), pipeline, proxyServerRouter)
	instance.Start()
}
