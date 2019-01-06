package main

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	pipeline := proxy.InitPipeline()
	instance := proxy.New("", os.Getenv("PORT"), pipeline, mux.NewRouter())
	instance.Start()
}
