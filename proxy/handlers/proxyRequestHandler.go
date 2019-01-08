package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/ASV44/DeliveryManagement-DS/proxy/util"
	"io"
	"net/http"
)

type ProxyRequestHandler struct {
	pipeline *proxy.Pipeline
}

func NewProxyRequestHandler(pipeline *proxy.Pipeline) *ProxyRequestHandler {
	return &ProxyRequestHandler{pipeline: pipeline}
}

func (handler *ProxyRequestHandler) forwardPostRequest(w http.ResponseWriter, req *http.Request, log string) {
	handler.pipeline.Log <- log + req.URL.Path
	url := proxy.WarehouseHost + req.URL.Path
	response, err := http.Post(url, req.Header.Get("Content-Type"), req.Body)

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.WarehouseRequestError, err)
		return
	}

	handler.pipeline.Log <- proxy.ForwardRequest + util.ToLog(req.ContentLength)
	length, err := io.Copy(w, response.Body)
	defer response.Body.Close()

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.WarehouseInvalidResponse, err)
		return
	}

	handler.pipeline.Log <- proxy.ForwardResponse + util.ToLog(length)
}
