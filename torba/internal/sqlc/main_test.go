package sqlc

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var test *Queries

func TestMain(m *testing.M) {

	ctx := context.Background()
	err := godotenv.Load("../../cmd/web/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("DB_STRING")

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("cannot ping db", err)
	}

	test = New(conn)

	os.Exit(m.Run())
}
