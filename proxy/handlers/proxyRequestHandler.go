package handlers

import (
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/ASV44/DeliveryManagement-DS/proxy/util"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProxyRequestHandler struct {
	pipeline *proxy.Pipeline
}

func NewProxyRequestHandler(pipeline *proxy.Pipeline) *ProxyRequestHandler {
	return &ProxyRequestHandler{pipeline: pipeline}
}

// Forward any request to new host, and returns http response or nil in case of failure.
// In case of failure handle error and sends to client json with error specification.
func (handler *ProxyRequestHandler) forwardRequest(w http.ResponseWriter, req *http.Request,
	url string, log string) *http.Response {
	handler.pipeline.Log <- log + req.URL.RequestURI()
	forwardUrlString := url + req.URL.RequestURI()
	forwardReq, _ := http.NewRequest(req.Method, forwardUrlString, req.Body)
	forwardReq.Header = req.Header
	response, err := http.DefaultClient.Do(forwardReq)

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.ForwardRequestError+forwardUrlString, err)
		return nil
	}

	handler.pipeline.Log <- proxy.ForwardRequest + util.ToLog(req.ContentLength)

	return response
}

// Forward response, in form of byte array, obtained from some source, to client.
func (handler *ProxyRequestHandler) forwardResponse(w http.ResponseWriter, response []byte) {
	length, err := w.Write(response)

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.DataSendFailed, err)
		return
	}

	handler.pipeline.Log <- proxy.ForwardResponse + strconv.Itoa(length)
}

// Forward request direct from proxy to warehouse,
// and automatically forward warehouse response to client
func (handler *ProxyRequestHandler) ForwardWarehouseReqRes(w http.ResponseWriter, req *http.Request, log string) {
	response := handler.forwardRequest(w, req, proxy.WarehouseHost, log)
	if response == nil {
		return
	}

	length, err := io.Copy(w, response.Body)
	defer response.Body.Close()

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.WarehouseInvalidResponse, err)
		return
	}

	handler.pipeline.Log <- proxy.ForwardResponse + util.ToLog(length)
}

// Forward request to warehouse and returns warehouse response, in form og byte array, to proxy
func (handler *ProxyRequestHandler) ForwardWarehouseReq(w http.ResponseWriter, req *http.Request, log string) []byte {
	response := handler.forwardRequest(w, req, proxy.WarehouseHost, log)
	if response == nil {
		return nil
	}
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		OnError(w, handler.pipeline, http.StatusInternalServerError, proxy.WarehouseInvalidResponse, err)
		return nil
	}

	return contents
}
