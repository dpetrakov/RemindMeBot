package bot

import (
	"log"

	"github.com/myuser/remindmebot/pkg/postgres"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	db    *postgres.DB
	nc    *nats.Conn
	tgBot *tgbotapi.BotAPI
}

func NewBot(db *postgres.DB, nc *nats.Conn) *Bot {
	// Load configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Connect to Telegram
	tgBot, err := tgbotapi.NewBotAPI(viper.GetString("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to connect to Telegram: %v", err)
	}

	return &Bot{
		db:    db,
		nc:    nc,
		tgBot: tgBot,
	}
}

func (b *Bot) HandleMessage(msg []byte) error {
	// Parse message
	text := string(msg)

	// Send message to Telegram
	msg := tgbotapi.NewMessage(viper.GetInt64("TELEGRAM_CHAT_ID"), text)
	_, err := b.tgBot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

