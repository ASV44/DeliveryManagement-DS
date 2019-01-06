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
	currentTime := time.Now()
	io.WriteString(w, server.WelcomeMessage)
	io.WriteString(w, currentTime.Format("2006-01-02 15:04:05"))
}

func (handler *ServerHandler) HandleError(w http.ResponseWriter, serverError models.ServerError) {
	w.WriteHeader(serverError.Status)
	jsonData, _ := json.Marshal(serverError)
	_, _ = w.Write(jsonData)
	handler.pipeline.Log <- fmt.Sprintf(server.SeverErrorLog, serverError.ClientErrorMessage, serverError.Error)
}

func (handler *ServerHandler) HandleInsertErrors(w http.ResponseWriter, insertErrors []models.InsertError) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonData, _ := json.Marshal(insertErrors)
	_, _ = w.Write(jsonData)
	handler.pipeline.Log <- server.OrdersRegisterFailed
	handler.pipeline.Log <- mappers.InsertErrorsToLog(insertErrors, server.OrdersRegisterFailedLog)
}
