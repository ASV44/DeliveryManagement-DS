package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/warehouse/server"
	"io"
	"net/http"
	"time"
)

type ServerHandler struct {
	pipeline *server.Pipeline
}

func NewServerHandler(pipeline *server.Pipeline) *ServerHandler {
	return &ServerHandler{
		pipeline: pipeline,
	}
}

func (handler *ServerHandler) RootHandler(w http.ResponseWriter, r *http.Request) {
	handler.pipeline.Log <- "request on '/' route"
	currentTime := time.Now()
	io.WriteString(w, "Delivery Management Distributed System ~DS\n")
	io.WriteString(w, currentTime.Format("2006-01-02 15:04:05"))
}