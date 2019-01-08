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

func HandleError(w http.ResponseWriter, pipeline *server.Pipeline, serverError models.ServerError) {
	w.WriteHeader(serverError.Status)
	jsonData, _ := json.Marshal(serverError)
	_, _ = w.Write(jsonData)
	pipeline.Log <- fmt.Sprintf(server.SeverErrorLog, serverError.ClientErrorMessage, serverError.Error)
}

func HandleOrdersErrors(w http.ResponseWriter, pipeline *server.Pipeline, orderErrors []models.OrderError,
	mainLog string, errorLog string) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonData, _ := json.Marshal(models.OrderErrors{Errors: orderErrors})
	_, _ = w.Write(jsonData)
	pipeline.Log <- mainLog
	pipeline.Log <- mappers.OrderErrorsToLog(orderErrors, errorLog)
}

func onError(w http.ResponseWriter, pipeline *server.Pipeline, status int, message string, err error) {
	e := models.ServerError{Status: status,
		ClientErrorMessage: message,
		Error:              err.Error()}
	HandleError(w, pipeline, e)
}

func onErrors(w http.ResponseWriter, pipeline *server.Pipeline, orderErrors []models.OrderError,
	mainLog string, errorLog string) {
	HandleOrdersErrors(w, pipeline, orderErrors, mainLog, errorLog)
}
