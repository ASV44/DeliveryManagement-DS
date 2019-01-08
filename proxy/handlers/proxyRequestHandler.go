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

func (handler *ProxyRequestHandler) forwardWarehouseRequest(w http.ResponseWriter, req *http.Request, log string) {
	handler.pipeline.Log <- log + req.URL.RequestURI()
	warehouseUrlString := proxy.WarehouseHost + req.URL.RequestURI()
	warehouseReq, _ := http.NewRequest(req.Method, warehouseUrlString, req.Body)
	warehouseReq.Header = req.Header
	response, err := http.DefaultClient.Do(warehouseReq)

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
