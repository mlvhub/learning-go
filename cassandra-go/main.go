package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

func PerformOperations() {
	cluster := gocql.NewCluster("127.0.0.1:29041", "127.0.0.1:29042", "127.0.0.1:29043")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "some_username",
		Password: "some_password",
	}

	cluster.Keyspace = "keyspace_name"

	cluster.Timeout = 5 * time.Second

	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Could not connect to Cassandra cluster: %v", err)
	}
	defer session.Close()

	keyspaceMeta, _ := session.KeyspaceMetadata("keyspace_name")
	log.Printf("KEYSPACE: %v", keyspaceMeta)

	if _, exists := keyspaceMeta.Tables["person"]; exists != true {
		log.Printf("Creating table...")
		if err := session.Query("CREATE TABLE keyspace_name.person (" +
			"id text, name text, phone text, " +
			"PRIMARY KEY (id))").Exec(); err != nil {
			log.Fatalf("Create table failed: %v", err)
		}
	}

	if err := session.Query("INSERT INTO keyspace_name.person (id, name, phone) VALUES (?, ?, ?)", "mlopez", "Miguel Lopez", "0101010101").Exec(); err != nil {
		log.Fatalf("Insert failed: %v", err)
	}

	var name string
	var phone string
	if err := session.Query("SELECT name, phone FROM keyspace_name.person WHERE id='mlopez'").Scan(&name, &phone); err != nil {
		if err != gocql.ErrNotFound {
			log.Fatalf("Query failed: %v", err)
		}
	}

	log.Printf("Name: %v", name)
	log.Printf("Phone: %v", phone)

	session.Query("DELETE id, name, phone FROM keyspace_name.person WHERE id='mlopez'").Exec()
}

func main() {
	PerformOperations()
}
