package models

type OrderError struct {
	Error   string `json:"error"`
	OrderID string `json:"id"`
}
