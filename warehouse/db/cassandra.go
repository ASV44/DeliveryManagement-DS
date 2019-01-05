package db

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
	"github.com/gocql/gocql"
	"time"
)

const (
	KEYSPACE = "delivery_management"
	HOST     = "cassandraSeed"
)

type Cassandra struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func (db *Cassandra) ConnectToCluster() {
	// connect to the cluster
	db.cluster = gocql.NewCluster(HOST)
	db.cluster.Port = 9042
	db.cluster.Keyspace = KEYSPACE
	db.cluster.Consistency = gocql.Quorum
	db.initSession()
}

func (db *Cassandra) initSession() {
	var err error
	db.session, err = db.cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		db.session.Close()
	} else {
		fmt.Println("Connected to Cassandra! Init done!")
	}
}

func (db *Cassandra) AddOrder(order models.Order) error {
	err := db.session.Query(
		`INSERT INTO orders(order_id, awb_number, allow_open_parcel,
								  created_date, labels, latitude, longitude,
								  service_payment, receiver_address,
								  receiver_address_locality, receiver_contact,
								  receiver_name, receiver_phone, shipper_address,
								  shipper_address_locality, shipper_contact,
								  shipper_name, shipper_phone, status_group_id,
								  today_important)
			   VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		order.Id, order.AwbNumber, order.AllowOpenParcel, order.CreatedDate,
		order.Labels, order.Latitude, order.Longitude, order.ServicePayment,
		order.ReceiverAddress, order.ReceiverAddressLocality, order.ReceiverContact,
		order.ReceiverName, order.ReceiverPhone, order.ShipperAddress,
		order.ShipperAddressLocality, order.ShipperContact,
		order.ShipperName, order.ShipperPhone, order.StatusGroupId,
		order.TodayImportant).Exec()

	return err
}

func (db *Cassandra) GetAllOrders() []models.Order {
	var orders []models.Order
	data := make(map[string]interface{})

	iter := db.session.Query("SELECT * FROM orders").Iter()
	for iter.MapScan(data) {
		orders = append(orders, models.Order{
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
		})
		data = make(map[string]interface{})
	}

	return orders
}
