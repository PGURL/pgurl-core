package cql

import (
	"time"

	"github.com/gocql/gocql"
)

var cqlClient *gocql.Session = dbcon()

func dbcon() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.DisableInitialHostLookup = true
	cluster.IgnorePeerAddr = true
	cluster.Timeout = 3 * time.Minute
	cluster.Keyspace = "short"
	cluster.Consistency = gocql.One
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 5}
	s, _ := cluster.CreateSession()
	return s
}
