package db

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/mappers"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
	"github.com/gocql/gocql"
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
	err := db.session.Query(InsertOrder, order.Id,
		order.AwbNumber,
		order.AllowOpenParcel,
		order.CreatedDate,
		order.Labels,
		order.Latitude,
		order.Longitude,
		order.ServicePayment,
		order.ReceiverAddress,
		order.ReceiverAddressLocality,
		order.ReceiverContact,
		order.ReceiverName,
		order.ReceiverPhone,
		order.ShipperAddress,
		order.ShipperAddressLocality,
		order.ShipperContact,
		order.ShipperName,
		order.ShipperPhone,
		order.StatusGroupId,
		order.TodayImportant).Exec()

	return err
}

func (db *Cassandra) RegisterNewOrders(order []models.Order) []models.OrderError {
	var errors []models.OrderError = nil
	for _, element := range order {
		err := db.AddOrder(element)
		if err != nil {
			errors = append(errors, models.OrderError{Error: err.Error(),
				OrderID: element.Id})
		}
	}

	return errors
}

func (db *Cassandra) getOrdersByQuery(query *gocql.Query) []models.Order {
	orders := make([]models.Order, 0)
	data := make(map[string]interface{})

	iter := query.Iter()
	for iter.MapScan(data) {
		orders = append(orders, mappers.DataToOrder(data))
		data = make(map[string]interface{})
	}

	return orders
}

func (db *Cassandra) GetAllOrders() []models.Order {
	query := db.session.Query(GetAllOrders)

	return db.getOrdersByQuery(query)
}

func (db *Cassandra) GetOrderById(id string) (models.Order, error) {
	var order models.Order
	data := make(map[string]interface{})
	err := db.session.Query(GetOrderById, id).MapScan(data)
	if len(data) != 0 {
		order = mappers.DataToOrder(data)
	}

	return order, err
}

func (db *Cassandra) GetOrdersByAWB(awbNumber string) []models.Order {
	query := db.session.Query(GetOrdersByAwb, awbNumber)

	return db.getOrdersByQuery(query)
}

func (db *Cassandra) UpdateOrderById(id string, values map[string]string) error {
	batch := db.session.NewBatch(gocql.LoggedBatch)
	for column, value := range values {
		query := SetUpdateColumn(UpdateOrderWithId, column)
		batch.Query(query, value, id)
	}
	err := db.session.ExecuteBatch(batch)

	return err
}

func (db *Cassandra) DeleteOrder(id string) error {
	err := db.session.Query(DeleteOrderWithId, id).Exec()

	return err
}

func (db *Cassandra) DeleteMultipleOrder(idList []string) []models.OrderError {
	var errors []models.OrderError = nil
	for _, id := range idList {
		err := db.DeleteOrder(id)
		if err != nil {
			errors = append(errors, models.OrderError{
				Error:   err.Error(),
				OrderID: id,
			})
		}
	}

	return errors
}
