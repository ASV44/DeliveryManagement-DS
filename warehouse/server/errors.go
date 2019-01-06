package server

const (
	SeverErrorLog           = "Error : %s \nCaused by : %s"
	IncorrectUrl            = "incorrect url! Url should have another parameters or queries for this request"
	InvalidJSONDecoding     = "Invalid json structure, decoding failed"
	InvalidJSONEncoding     = "Json encoding failed"
	DataSendFailed          = "Error while sending data"
	OrderAddFailed          = "Error while adding order : "
	OrderWithIdNotFound     = "Order with requested Id Not found : "
	OrdersRegisterFailed    = "Insertion of orders failed : "
	OrdersRegisterFailedLog = "Order %s insertion failed. Reason %s"
	OrderUpdateFailed       = "Update of order failed : "
	OrderDeleteFailed       = "Delete of order failed : "
	OrdersIdNotPresent      = "Orders ID for query not present in url, it must include '?id={order_id}'"
	OrdersDeleteFailedLog   = "Order %s delete failed. Reason %s"
)
