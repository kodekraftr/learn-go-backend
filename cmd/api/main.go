package main

import (
	"kodekraftr/learn-go-backend/internal/db"
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
	connStr := os.Getenv("DATABASE_URL")

	config := config{
		addr: addr,
		db: dbConfig{
			connStr:      connStr,
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
	}

	db, err := db.New(
		config.db.connStr,
		config.db.maxOpenConns,
		config.db.maxIdleConns,
		config.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection established")
	store := store.NewStorage(db)

	app := &application{
		config: config,
		store:  store,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))

}
