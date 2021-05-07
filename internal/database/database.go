package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/kisstupid/go-gin/configs"
)

var db *pgx.Conn

func Init() {
	config := configs.GetConfig()
	database, err := pgx.Connect(context.Background(), config.GetString("db.url"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db = database
}

func GetDB() *pgx.Conn {
	return db
}
