package reminder

import (
	"log"

	"github.com/dpetrakov/remindmebot/pkg/bot"
	"github.com/dpetrakov/remindmebot/pkg/nats"
	"github.com/dpetrakov/remindmebot/pkg/postgres"
)

type ReminderService struct {
	db *postgres.DB
	nc *nats.Conn
}

func NewReminderService(db *postgres.DB, nc *nats.Conn) *ReminderService {
	return &ReminderService{db, nc}
}

func (s *ReminderService) Start() error {
	// Start bot service
	if err := bot.Start(s.db, s.nc); err != nil {
		return err
	}

	// Start reminder service
	if err := bot.StartReminders(s.db, s.nc); err != nil {
		return err
	}

	// Subscribe to "reminder_added" queue
	if _, err := s.nc.Subscribe("reminder_added", func(msg *nats.Msg) {
		reminder := &bot.ReminderEvent{}
		if err := reminder.FromJSON(msg.Data); err != nil {
			log.Printf("Failed to parse reminder event: %v", err)
			return
		}

		if err := bot.HandleReminderAdded(s.db, s.nc, reminder); err != nil {
			log.Printf("Failed to handle reminder event: %v", err)
		}
	}); err != nil {
		return err
	}

	return nil
}

