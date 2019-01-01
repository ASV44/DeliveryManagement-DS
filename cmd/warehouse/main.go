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
	serverRouter := router.New(handlers.NewServerHandler(pipeline), handlers.NewOrdersHandler(cassandra, pipeline))
	instance := server.New("", os.Getenv("PORT"), pipeline, serverRouter, cassandra)
	instance.Start()
}
