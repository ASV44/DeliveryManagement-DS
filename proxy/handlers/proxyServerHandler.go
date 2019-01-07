package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"io"
	"net/http"
	"time"
)

type ProxyServerHandler struct {
	pipeline *proxy.Pipeline
}

func NewProxyServerHandler(pipeline *proxy.Pipeline) *ProxyServerHandler {
	return &ProxyServerHandler{
		pipeline: pipeline,
	}
}

func (handler *ProxyServerHandler) ProxyRootHandler(w http.ResponseWriter, r *http.Request) {
	handler.pipeline.Log <- proxy.GetRequestToRoot
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(w, proxy.WelcomeMessage+proxy.TimeMessage+currentTime)
}
