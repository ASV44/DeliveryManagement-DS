package models

type ProxyServerError struct {
	Status             int    `json:"status"`
	ClientErrorMessage string `json:"serverMessage"`
	Error              string `json:"error"`
}
