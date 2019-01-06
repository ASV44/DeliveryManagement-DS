package models

type InsertError struct {
	Error          string `json:"error"`
	OrderID        string `json:"id"`
	OrderAwbNumber string `json:"awbNumber"`
}
