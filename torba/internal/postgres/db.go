package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type DB struct {
	pdb   *pgxpool.Pool
	infol *log.Logger
	errl  *log.Logger
}

func Init(ctx context.Context) (*DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	c, err := pgxpool.ParseConfig(os.Getenv("DB_STRING"))
	if err != nil {
		return nil, errors.New("error parsing config by pgxpool")
	}

	c.MaxConns = 25
	c.MinConns = 5
	c.MaxConnIdleTime = 5 * time.Minute
	c.MaxConnLifetime = 25 * time.Minute
	c.HealthCheckPeriod = 1 * time.Minute
	c.ConnConfig.ConnectTimeout = 5 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, c)
	if err != nil {
		return nil, errors.New("error connecting to database")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.New("error pinging database")
	}

	log.Println("connected to database")

	row := pool.QueryRow(context.Background(), "SELECT now()")

	var currentTime time.Time

	err = row.Scan(&currentTime)
	if err != nil {
		return nil, errors.New("error getting current time")
	}
	log.Println("current time:", currentTime)

	pdb := DB{pdb: pool, infol: log.New(os.Stdout, "INFO [DB:]\t", log.Ldate|log.Ltime|log.Lshortfile), errl: log.New(os.Stdout, "ERROR [DB:]\t", log.Ldate|log.Ltime|log.Lshortfile)}

	return &pdb, nil
}
