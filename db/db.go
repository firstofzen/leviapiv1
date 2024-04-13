package db

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func Connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Consistency = gocql.LocalQuorum
	if sess, err := gocqlx.WrapSession(cluster.CreateSession()); err != nil {
		fmt.Print("connecting to db...")
		return gocqlx.Session{}, errors.New(err.Error())
	} else {
		return sess, nil
	}
}
