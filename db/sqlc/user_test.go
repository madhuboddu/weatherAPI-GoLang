package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func TestCreateRecord(t *testing.T) {
	arg := CreateRecordParams{

		Name:        "Madhu",
		Zipcode:     "N2J 2C1",
		Location:    "waterloo",
		Temperature: 18.0,
		FeelsLike:   16.0,
	}

	rec, err := testQueries.CreateRecord(context.Background(), arg)

	require.Nil(t, err)
	require.Equal(t, arg.Name, rec.Name)
}
