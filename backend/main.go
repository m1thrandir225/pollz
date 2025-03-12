package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"log"
	"m1thrandir225/cicd2025/api"
	db "m1thrandir225/cicd2025/db/sqlc"
	"m1thrandir225/cicd2025/util"
)

type ginServerConfig struct {
	Config util.Config
	Store  db.Store
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err.Error())
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err.Error())
	}

	connPool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	store := db.NewStore(connPool)

	serverConfig := ginServerConfig{
		Config: config,
		Store:  store,
	}

	runGinServer(serverConfig)
}

func runGinServer(settings ginServerConfig) {
	server, err := api.NewServer(settings.Store, settings.Config)
	if err != nil {
		log.Fatal("Cannot create server")
	}
	err = server.Start(settings.Config.HTTPServerAddress)

	if err != nil {
		log.Fatal("Cannot start server")
	}
}
