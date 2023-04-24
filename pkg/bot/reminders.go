package bot

import (
	"log"
	"time"

	"github.com/myuser/remindmebot/pkg/postgres"

	"github.com/nats-io/nats.go"
)

func StartReminders(db *postgres.DB, nc *nats.Conn) error {
	// Get all reminders from the database
	reminders, err := db.GetAllReminders()
	if err != nil {
		return err
	}

	// Schedule reminders
	for _, reminder := range reminders {
		if err := HandleRemindEvent(db, nc, reminder); err != nil {
			log.Printf("Failed to schedule reminder event: %v", err)
		}
	}

	return nil
}

func HandleReminderAdded(db *postgres.DB, nc *nats.Conn, reminder *ReminderEvent) error {
	// Schedule reminder event
	if err := HandleRemindEvent(db, nc, reminder); err != nil {
		return err
	}

	return nil
}

