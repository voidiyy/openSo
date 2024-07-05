package main

import (
	"github.com/joho/godotenv"
	"iblan/cmd/api"
	"iblan/cmd/storage"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("LISTEN_ADDR")

	storage, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	storage.Migrate()
	userSrv := api.NewAPIServer(addr, storage)

	userSrv.Run()

}
