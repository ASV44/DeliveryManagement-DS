package models

type ServerError struct {
	Status             int    `json:"status"`
	ClientErrorMessage string `json:"serverMessage"`
	Error              string `json:"error"`
}
