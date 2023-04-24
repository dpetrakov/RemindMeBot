package reminder

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dpetrakov/remindmebot/pkg/bot"
	"github.com/dpetrakov/remindmebot/pkg/postgres"
)

type PostgresReminderStore struct {
	db *sql.DB
}

func NewPostgresReminderStore(db *sql.DB) *PostgresReminderStore {
	return &PostgresReminderStore{db}
}

func (s *PostgresReminderStore) CreateReminder(reminder *bot.ReminderEvent) error {
	query := `INSERT INTO reminders (user_id, reminder_time, reminder_text) VALUES ($1, $2, $3) RETURNING id`
	row := s.db.QueryRow(query, reminder.UserID, reminder.ReminderTime, reminder.ReminderText)

	var id int64
	if err := row.Scan(&id); err != nil {
		return err
	}

	reminder.ID = id

	return nil
}

func (s *PostgresReminderStore) GetAllReminders() ([]*bot.ReminderEvent, error) {
	query := `SELECT * FROM reminders`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reminders := []*bot.ReminderEvent{}
	for rows.Next() {
		reminder := &bot.ReminderEvent{}
		if err := rows.Scan(&reminder.ID, &reminder.UserID, &reminder.ReminderTime, &reminder.ReminderText); err != nil {
			return nil, err
		}
		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func (s *PostgresReminderStore) GetRemindersByUserID(userID int64) ([]*bot.ReminderEvent, error) {
	query := `SELECT * FROM reminders WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reminders := []*bot.ReminderEvent{}
	for rows.Next() {
		reminder := &bot.ReminderEvent{}
		if err := rows.Scan(&reminder.ID, &reminder.UserID, &reminder.ReminderTime, &reminder.ReminderText); err != nil {
			return nil, err
		}
		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func (s *PostgresReminderStore) DeleteReminder(id int64) error {
	query := `DELETE FROM reminders WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresReminderStore) DeleteRemindersByUserID(userID int64) error {
	query := `DELETE FROM reminders WHERE user_id = $1`
	_, err := s.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresReminderStore) Migrate() error {
	// Create reminders table
	query := `
		CREATE TABLE IF NOT EXISTS reminders (
			id SERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			reminder_time TIMESTAMP NOT NULL,
			reminder_text TEXT NOT NULL
		);
	`
	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("failed to create reminders table: %v", err)
	}

	return nil
}

func (s *PostgresReminderStore) Close() error {
	return s.db.Close()
}

func (s *PostgresReminderStore) AutoMigrate() error {
	// Migrate

