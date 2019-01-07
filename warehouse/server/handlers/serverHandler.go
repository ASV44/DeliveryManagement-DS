package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/mappers"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
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
	handler.pipeline.Log <- server.GetRequestToRoot
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(w, server.WelcomeMessage+server.TimeMessage+currentTime)
}

func (handler *ServerHandler) HandleError(w http.ResponseWriter, serverError models.ServerError) {
	w.WriteHeader(serverError.Status)
	jsonData, _ := json.Marshal(serverError)
	_, _ = w.Write(jsonData)
	handler.pipeline.Log <- fmt.Sprintf(server.SeverErrorLog, serverError.ClientErrorMessage, serverError.Error)
}

func (handler *ServerHandler) HandleOrdersErrors(w http.ResponseWriter, orderErrors []models.OrderError,
	mainLog string, errorLog string) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonData, _ := json.Marshal(models.OrderErrors{Errors: orderErrors})
	_, _ = w.Write(jsonData)
	handler.pipeline.Log <- mainLog
	handler.pipeline.Log <- mappers.OrderErrorsToLog(orderErrors, errorLog)
}
