package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func TestCreateRecord(t *testing.T) {
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

	if testQueries == nil {
		t.Fatal("testQueries is not initialized")
	}

	arg := CreateRecordParams{

		Name:        "Madhu",
		Zipcode:     "N2J 2C1",
		Location:    "waterloo",
		Temperature: 18.0,
		FeelsLike:   16.0,
	}

	rec, err := testQueries.CreateRecord(context.Background(), arg)

	if err != nil {
		t.Fatalf("failed to create record: %v", err)
	}

	require.Nil(t, err)
	require.Equal(t, arg.Name, rec.Name)
}
