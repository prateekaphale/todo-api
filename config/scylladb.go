package config

import (
    "log"

    "github.com/gocql/gocql"
)

var Session *gocql.Session

func InitScylla() {
    cluster := gocql.NewCluster("localhost")
    cluster.Keyspace = "todo"
    cluster.Consistency = gocql.Quorum

    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }

    Session = session
}
