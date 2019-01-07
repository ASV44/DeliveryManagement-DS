package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/proxy"
	"github.com/ASV44/DeliveryManagement-DS/proxy/models"
	"io"
	"io/ioutil"
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
	currentTime := time.Now().Format("2006-01-02 15:04:05\n")
	proxyMessage := proxy.WelcomeMessage + proxy.TimeMessage + currentTime
	response, err := http.Get(proxy.WarehouseHost)

	if err != nil {
		handler.OnError(w, http.StatusInternalServerError, proxy.WarehouseRequestError, err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		handler.OnError(w, http.StatusInternalServerError, proxy.WarehouseInvalidResponse, err)
		return
	}

	warehouseMessage := string(data)
	io.WriteString(w, proxyMessage+warehouseMessage)
}

func HandleError(w http.ResponseWriter, pipeline *proxy.Pipeline, serverError models.ProxyServerError) {
	w.WriteHeader(serverError.Status)
	jsonData, _ := json.Marshal(serverError)
	_, _ = w.Write(jsonData)
	pipeline.Log <- fmt.Sprintf(proxy.SeverErrorLog, serverError.ClientErrorMessage, serverError.Error)
}

func (handler *ProxyServerHandler) OnError(w http.ResponseWriter, status int, message string, err error) {
	serverError := models.ProxyServerError{Status: status,
		ClientErrorMessage: message,
		Error:              err.Error()}
	HandleError(w, handler.pipeline, serverError)
}
