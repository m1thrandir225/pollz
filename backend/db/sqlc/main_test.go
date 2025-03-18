package db

import (
	"context"
	"log"
	"m1thrandir225/cicd2025/util"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.TestingDbSource)
	if err != nil {
		log.Fatal("cannot connect to test database:", err)
	}
	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
