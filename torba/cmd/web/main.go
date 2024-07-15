package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
	"torba/internal/sqlc"
)

type application struct {
	infoL  *log.Logger
	errorL *log.Logger
	db     *sqlc.Queries
}

func NewApplication(store *sqlc.Queries) *application {
	return &application{
		infoL:  log.New(os.Stdout, "INFO[:]\t", log.Ldate|log.Ltime),
		errorL: log.New(os.Stderr, "ERROR[:]\t", log.Ldate|log.Ltime),
		db:     store,
	}
}

func run() (*sqlc.Queries, error) {
	ctx := context.Background()

	godotenv.Load()
	//url := os.Getenv("DB_URL")
	str := os.Getenv("DB_STRING")

	conn, err := pgx.Connect(ctx, str)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	queries := sqlc.New(conn)

	var users []sqlc.User
	users, err = queries.ListUserName(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	log.Printf("created user: %s,,, ", users)

	return queries, nil
}

func main() {

	postgr, err := run()
	if err != nil {
		log.Fatal(err)
	}

	app := NewApplication(postgr)

	app.infoL.Printf("database connection established: %v", postgr)

	//configs

	err = godotenv.Load()
	if err != nil {
		app.infoL.Fatal("Error loading .env file")
	}

	addr := os.Getenv("LISTEN_ADDR")

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app.infoL.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	app.errorL.Fatal(err)
}
