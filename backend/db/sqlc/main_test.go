package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	connString := os.Getenv("TESTING_DB_SOURCE")
	if connString == "" {
		log.Fatal("TESTING_DB_SOURCE environment variable not set")
	}

	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to test database:", err)
	}
	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
