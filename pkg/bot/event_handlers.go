package bot

import (
	"log"
	"time"

	"github.com/dpetrakov/remindmebot/pkg/postgres"

	"github.com/nats-io/nats.go"
)

func HandleRemindEvent(db *postgres.DB, nc *nats.Conn, event *ReminderEvent) error {
	// Calculate time difference between now and the reminder time
	duration := event.ReminderTime.Sub(time.Now())
	if duration <= 0 {
		return nil
	}

	// Wait until reminder time
	time.Sleep(duration)

	// Get chat ID from the database
	chatID, err := db.GetChatID(event.UserID)
	if err != nil {
		return err
	}

	// Send reminder message to the bot
	msg := []byte(event.ReminderText)
	if err := nc.Publish("remindme", msg); err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	return nil
}

