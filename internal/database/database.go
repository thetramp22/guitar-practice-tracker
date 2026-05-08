package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/thetramp22/rifflog/internal/config"
)

func NewConnection() *pgx.Conn {
	dbConfig, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	var conn *pgx.Conn

	for i := 0; i < 10; i++ {
		conn, err = pgx.ConnectConfig(context.Background(), dbConfig)

		if err == nil {
			log.Println("Connected to PostgreSQL")
			return conn
		}

		log.Printf("Database not ready... retrying (%d/10)\n", i+1)

		time.Sleep(2 * time.Second)
	}

	log.Fatal("Unable to connect to database:", err)

	return nil
}
