package main

import (
	"github.com/ASV44/DeliveryManagement-DS/warehouse/db"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server/handlers"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server/router"
)
import "os"

func main() {
	cassandra := &db.Cassandra{}
	cassandra.ConnectToCluster()
	pipeline := server.InitPipeline()
	serverHandler := handlers.NewServerHandler(pipeline)
	ordersHandler := handlers.NewOrdersHandler(serverHandler, cassandra, pipeline)
	serverRouter := router.New(serverHandler, ordersHandler)
	instance := server.New("", os.Getenv("PORT"), pipeline, serverRouter, cassandra)
	instance.Start()
}
