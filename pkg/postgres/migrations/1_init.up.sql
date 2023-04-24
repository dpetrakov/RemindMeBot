CREATE TABLE IF NOT EXISTS reminders (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    reminder_time TIMESTAMP NOT NULL,
    reminder_text TEXT NOT NULL
);

