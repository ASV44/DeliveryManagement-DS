package models

import "time"

type Order struct {
	Id                      string    `json:"id"`
	AwbNumber               string    `json:"awbNumber"`
	AllowOpenParcel         bool      `json:"allowOpenParcel"`
	CreatedDate             time.Time `json:"createdDate"`
	Labels                  []string  `json:"labels"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	ServicePayment          float64   `json:"servicePayment"`
	ReceiverAddress         string    `json:"receiverAddress"`
	ReceiverAddressLocality string    `json:"receiverAddressLocality"`
	ReceiverContact         string    `json:"receiverContact"`
	ReceiverName            string    `json:"receiverName"`
	ReceiverPhone           string    `json:"receiverPhone"`
	ShipperAddress          string    `json:"shipperAddress"`
	ShipperAddressLocality  string    `json:"shipperAddressLocality"`
	ShipperContact          string    `json:"shipperContact"`
	ShipperName             string    `json:"shipperName"`
	ShipperPhone            string    `json:"shipperPhone"`
	StatusGroupId           int       `json:"statusGroupId"`
	TodayImportant          bool      `json:"todayImportant"`
}
