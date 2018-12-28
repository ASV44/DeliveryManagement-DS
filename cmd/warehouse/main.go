package main

import (
	"github.com/ASV44/DeliveryManagement-DS/db"
	"github.com/ASV44/DeliveryManagement-DS/server"
)
import "os"

func main() {
	cassandra := &db.Cassandra{}
	cassandra.ConnectToCluster()
	instance := server.New("", os.Getenv("PORT"), cassandra)
	instance.Start()
}
