package config

import (
    "github.com/gocql/gocql"
)

var Session *gocql.Session

func InitScylla() {
    cluster := gocql.NewCluster("172.17.0.2")
    cluster.Keyspace = "your_keyspace_name"
    cluster.Consistency = gocql.Quorum
    var err error
    Session, err = cluster.CreateSession()
    if err != nil {
        panic(err)
    }
}
