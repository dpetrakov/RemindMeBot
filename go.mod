module github.com/dpetrakov/remindmebot

go 1.17

require (
	github.com/lib/pq v1.10.2
	github.com/nats-io/nats.go v1.25.0
	github.com/spf13/viper v1.15.0
	gopkg.in/telegram-bot-api.v4 v4.6.4
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nats-io/nats-server/v2 v2.9.16 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dpetrakov/remindmebot/cmd/bot => ./cmd/bot

replace github.com/dpetrakov/remindmebot/cmd/reminder => ./cmd/reminder

replace github.com/dpetrakov/remindmebot/pkg/bot => ./pkg/bot

replace github.com/dpetrakov/remindmebot/pkg/nats => ./pkg/nats

replace github.com/dpetrakov/remindmebot/pkg/postgres => ./pkg/postgres

replace github.com/dpetrakov/remindmebot/pkg/reminder => ./pkg/reminder
