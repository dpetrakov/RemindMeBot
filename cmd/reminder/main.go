package main

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"

	"github.com/dpetrakov/remindmebot/pkg/postgres"
	"github.com/dpetrakov/remindmebot/pkg/reminder"
)

func main() {
	// Load configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Connect to Postgres
	db, err := postgres.ConnectDB(viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connect to Nats
	nc, err := nats.Connect(viper.GetString("NATS_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to Nats: %v", err)
	}
	defer nc.Close()

	// Create reminder
	r := reminder.NewReminder(db, nc)

	// Run reminder loop
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := r.Remind()
			if err != nil {
				log.Printf("Failed to send reminders: %v", err)
			}
		}
	}
}

