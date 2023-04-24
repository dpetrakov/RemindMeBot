module github.com/dpetrakov/remindmebot

go 1.17

require (
    github.com/go-co-op/gocron v0.4.2
    github.com/lib/pq v1.10.2
    github.com/nats-io/nats.go v1.10.3
    github.com/robfig/cron/v3 v3.0.1
)

replace github.com/dpetrakov/remindmebot/cmd/bot => ./cmd/bot
replace github.com/dpetrakov/remindmebot/cmd/reminder => ./cmd/reminder
replace github.com/dpetrakov/remindmebot/pkg/bot => ./pkg/bot
replace github.com/dpetrakov/remindmebot/pkg/nats => ./pkg/nats
replace github.com/dpetrakov/remindmebot/pkg/postgres => ./pkg/postgres
replace github.com/dpetrakov/remindmebot/pkg/reminder => ./pkg/reminder

