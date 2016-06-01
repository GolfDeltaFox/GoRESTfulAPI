package database

import (
	"github.com/gocql/gocql"
	"fmt"
	"log"
)

var(
	Cassandra *gocql.Session
)

type Database struct {
	Hosts    []string
	Keyspace string

}


// Connect to the database
func (d Database) Connect() {
	var err error
	cluster := gocql.NewCluster(d.Hosts...)
	cluster.Keyspace = d.Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	if Cassandra, err = cluster.CreateSession(); err != nil {
		log.Println("Cassandra Driver Error", err)
	}
	defer Cassandra.Close()
}

func (d Database) request(table string) string{
	return fmt.Sprintf(`CREATE TABLE %s.%s (
	id	uuid,
	email	varchar,
	createdAt	timestamp,
	verified	boolean,
	PRIMARY KEY (email)
	)`,d.Keyspace, table)
}

func (d Database) Init() {
	Cassandra.Query( "DESCRIBE KEYSPACES")
	//if err := Cassandra.Query( d.request("user")); err != nil {
	//	log.Fatal("create:", err)
	//}
	println("ok")
}
