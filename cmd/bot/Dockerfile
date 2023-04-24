# Начальный образ, на основе которого будем строить свой образ
FROM golang:1.16-alpine

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копирование необходимых файлов в контейнер
COPY cmd/bot/main.go .
COPY pkg/bot/bot.go ./pkg/bot/
COPY pkg/bot/event_handlers.go ./pkg/bot/
COPY pkg/bot/message_handlers.go ./pkg/bot/
COPY pkg/bot/reminders.go ./pkg/bot/
COPY pkg/postgres/postgres.go ./pkg/postgres/
COPY pkg/nats/nats.go ./pkg/nats/

# Установка зависимостей
RUN apk add --no-cache git
RUN go mod init bot && go mod tidy
RUN go get -d -v ./...
RUN go build -o /app/bot

# Открываем порт, на котором будет работать приложение
EXPOSE 8080

# Запуск приложения
CMD ["/app/bot"]
