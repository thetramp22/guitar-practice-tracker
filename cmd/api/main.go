package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/thetramp22/rifflog/internal/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbConfig, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Unable to parse DB config:", err)
	}

	var conn *pgx.Conn

	for i := 0; i < 10; i++ {
		conn, err = pgx.ConnectConfig(context.Background(), dbConfig)

		if err == nil {
			log.Println("Connected to PostgreSQL")
			break
		}

		log.Printf("Database not ready yet... retrying (%d/10)\n", i+1)

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	log.Println("Connected to PostgreSQL")

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
