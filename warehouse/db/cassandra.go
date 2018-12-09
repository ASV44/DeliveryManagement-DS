package db

import (
	"fmt"
	"github.com/gocql/gocql"
)

const (
	KEYSPACE = "DeliveryManagement"
	HOST 	 = "cassandraSeed"
	REPLICATION_FACTOR = 3
)

type Cassandra struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func (db *Cassandra) ConnectToCluster() {
	// connect to the cluster
	db.cluster = gocql.NewCluster(HOST)
	db.cluster.Port = 9042
	db.cluster.Consistency = gocql.Quorum
	db.session, _ = db.cluster.CreateSession()
	defer db.session.Close()

	db.initSession()
}

func (db *Cassandra) initSession() {
	var err error
	if db.session == nil || db.session.Closed() {
		db.session, err = db.cluster.CreateSession()
	}
	if err != nil {
		fmt.Println(err)
	} else {
		db.initKeySpace()
		db.initTables()
		fmt.Println("Connected to Cassandra! Init done!")
	}

}

func (db *Cassandra) initKeySpace() {
	err := db.session.Query(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s 
	WITH replication = {
		'class' : 'SimpleStrategy',
		'replication_factor' : %d
	}`, KEYSPACE, REPLICATION_FACTOR)).Exec()

	if err != nil {
		fmt.Println("Error while creating keyspace")
		fmt.Println(err)
	}
}

func (db *Cassandra) initTables() {
	err := db.session.Query(`CREATE TABLE IF NOT EXISTS orders(
		order_id text,
		awb_number text,
		allow_open_parcel text,
		created_date timestamp,
		labels list<text>,
		latitude float,
		longitude float,
		service_payment float,
		receiver_address text,
		receiver_address_locality text,
		receiver_contact text,
		receiver_name text,
		receiver_phone text,
		shipper_address text,
		shipper_address_locality text,
		shipper_contact text,
		shipper_name text,
		shipper_phone text,
		status_group_id int,
		todayImportant boolean,
		PRIMARY KEY ((order_id, awb_number), created_date))`).Exec()

	if err != nil {
		fmt.Println("Error while creating tables")
		fmt.Println(err)
	}
}