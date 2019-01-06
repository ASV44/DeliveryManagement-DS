package server

const (
	SeverErrorLog           = "Error : %s \nCaused by : %s"
	InvalidJSONDecoding     = "Invalid json structure, decoding failed"
	InvalidJSONEncoding     = "Json encoding failed"
	DataSendFailed          = "Error while sending data"
	OrderAddFailed          = "Error while adding order : "
	OrderWithIdNotFound     = "Order with requested Id Not found : "
	OrdersRegisterFailed    = "Insertion of orders failed : "
	OrdersRegisterFailedLog = "Order %s / %s insertion failed. Reason %s"
	OrderUpdateFailed       = "Update of order failed : "
)
