package mappers

import (
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
	"time"
)

func DataToOrder(data map[string]interface{}) models.Order {
	return models.Order{
		Id:                      data["order_id"].(string),
		AwbNumber:               data["awb_number"].(string),
		AllowOpenParcel:         data["allow_open_parcel"].(bool),
		CreatedDate:             data["created_date"].(time.Time),
		Labels:                  data["labels"].([]string),
		Latitude:                data["latitude"].(float64),
		Longitude:               data["longitude"].(float64),
		ServicePayment:          data["service_payment"].(float64),
		ReceiverAddress:         data["receiver_address"].(string),
		ReceiverAddressLocality: data["receiver_address_locality"].(string),
		ReceiverContact:         data["receiver_contact"].(string),
		ReceiverName:            data["receiver_name"].(string),
		ReceiverPhone:           data["receiver_phone"].(string),
		ShipperAddress:          data["shipper_address"].(string),
		ShipperAddressLocality:  data["shipper_address_locality"].(string),
		ShipperContact:          data["shipper_contact"].(string),
		ShipperName:             data["shipper_name"].(string),
		ShipperPhone:            data["shipper_phone"].(string),
		StatusGroupId:           data["status_group_id"].(int),
		TodayImportant:          data["today_important"].(bool),
	}
}
