package main

import (
	"fmt"
	"log"
	"time"
	"tux21b.org/v1/gocql"
	"tux21b.org/v1/gocql/uuid"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "skeezy"
	cluster.Consistency = gocql.Quorum

	cass, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Error creating Cassandra session: %v", err))
	}
	defer cass.Close()

	authors := []string{"spammer", "program"}
	body := "blah blah blah keyword blah"
	created := time.Now()
	tags := []string{"spam", "test", "keywords"}

	for i := 0; i < 10000; i++ {
		id := uuid.TimeUUID()
		fmt.Printf("id: %s\n", id)
		err = cass.Query(`INSERT INTO posts (id, authors, body, dttm, tags) VALUES (?, ?, ?, ?, ?)`, id, authors, body, created, tags).Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
}
