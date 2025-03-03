package main

import (
	"kodekraftr/learn-go-backend/internal/store"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("ADDR")

	config := config{
		addr: addr,
	}

	store := store.NewStorage(nil)

	app := &application{
		config: config,
		store:  store,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))

}
