package db

import (
	"fmt"
	"github.com/gocql/gocql"
)

const (
	KEYSPACE = "DeliveryManagement"
	HOST 	 = "cassandraSeed"
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
		fmt.Println("Connected to Cassandra! Init done!")
	}

}