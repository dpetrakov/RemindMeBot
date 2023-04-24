package bot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dpetrakov/remindmebot/pkg/postgres"

	"github.com/nats-io/nats.go"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleStartMessage(db *postgres.DB, nc *nats.Conn, chatID int64, message string) error {
	// Save user to the database
	userID, err := db.CreateUser(chatID)
	if err != nil {
		return err
	}

	// Send welcome message to user
	msg := tgbotapi.NewMessage(chatID, "Welcome to RemindMeBot!")
	if _, err := b.tgBot.Send(msg); err != nil {
		return err
	}

	return nil
}

func HandleReminderMessage(db *postgres.DB, nc *nats.Conn, chatID int64, message string) error {
	// Parse reminder text and time
	parts := strings.Split(message, " ")
	if len(parts) != 3 {
		return fmt.Errorf("invalid reminder format")
	}

	reminderTime, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return fmt.Errorf("invalid time format: %v", err)
	}

	reminderText := parts[1]

	// Create reminder event
	event := &ReminderEvent{
		UserID:       chatID,
		ReminderTime: reminderTime,
		ReminderText: reminderText,
	}

	// Save reminder to the database
	if err := db.CreateReminder(event); err != nil {
		return err
	}

	// Schedule reminder event
	if err := HandleRemindEvent(db, nc, event); err != nil {
		return err
	}

	return nil
}

