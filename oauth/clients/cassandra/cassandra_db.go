package cassandra

import (
	"github.com/gocql/gocql"
	_ "github.com/gocql/gocql"
)

const (
	Keyspace = "oauth"
)

var (
	session *gocql.Session
)

func init() {
	// https://github.com/gocql/gocql#example
	// connect to the cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = Keyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
