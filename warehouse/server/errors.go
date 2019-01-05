package server

const (
	SeverErrorLog       = "Error : %s \nCaused by : %s"
	InvalidJSONDecoding = "Invalid json structure, decoding failed"
	InvalidJSONEncoding = "Json encoding failed"
	DataSendFailed      = "Error while sending data"
	OrderAddFailed      = "Error while adding order : "
	OrderWithIdNotFound = "Order with requested Id Not found : "
)
