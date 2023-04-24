package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"

	"github.com/myuser/remindmebot/pkg/bot"
	"github.com/myuser/remindmebot/pkg/postgres"
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

	// Create bot
	b := bot.NewBot(db, nc)

	// Subscribe to messages
	_, err = nc.Subscribe("remindme", func(m *nats.Msg) {
		err := b.HandleMessage(m.Data)
		if err != nil {
			log.Printf("Failed to handle message: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to messages: %v", err)
	}

	fmt.Println("Bot is running...")
	fmt.Println("Press CTRL-C to stop")

	// Wait forever
	select {}
}

