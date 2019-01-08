package main

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/ASV44/DeliveryManagement-DS/proxy/caching"
	"github.com/ASV44/DeliveryManagement-DS/proxy/handlers"
	"github.com/ASV44/DeliveryManagement-DS/proxy/router"
	"os"
)

func main() {
	redisClient := caching.InitRedis()
	pipeline := proxy.InitPipeline()
	serverHandler := handlers.NewProxyServerHandler(pipeline)
	proxyRequestHandler := handlers.NewProxyRequestHandler(pipeline)
	ordersProxyHandler := handlers.NewOrdersProxyHandler(proxyRequestHandler, pipeline, redisClient)
	proxyServerRouter := router.NewProxyRouter(serverHandler, ordersProxyHandler)
	instance := proxy.New("", os.Getenv("PORT"), pipeline, proxyServerRouter)
	instance.Start()
}
