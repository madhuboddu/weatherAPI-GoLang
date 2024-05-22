package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgresql://admin:secret@localhost:5431/weatherdb?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}
	fmt.Println("Initilize test queries : ", testQueries)
	testQueries = New(conn) // Make sure this matches how you've structured your Queries constructor
	os.Exit(m.Run())
}
