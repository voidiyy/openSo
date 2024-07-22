package main

import (
	"context"
	"log"
	"net/http"
	"openSo/internal/postgres"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Application struct {
	addr   string
	infoL  *log.Logger
	errorL *log.Logger
	db     *postgres.DB
}

func NewApplication(pool *postgres.DB) *Application {
	return &Application{
		addr:   os.Getenv("LISTEN_ADDR"),
		infoL:  log.New(os.Stdout, "INFO [:]\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorL: log.New(os.Stderr, "ERROR[:]\t", log.Ldate|log.Ltime|log.Lshortfile),
		db:     pool,
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pdb, err := postgres.Init(ctx)
	if err != nil {
		log.Printf("failed to initialize database: %v", err)
	}

	app := NewApplication(pdb)

	err = godotenv.Load()
	if err != nil {
		app.infoL.Fatal("Error loading .env file")
	}

	srv := &http.Server{
		Addr:         app.addr,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app.infoL.Printf("Starting server on %s", app.addr)
	err = srv.ListenAndServe()
	app.errorL.Fatal(err)
}
