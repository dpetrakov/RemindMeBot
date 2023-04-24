# Базовый образ для Golang
FROM golang:1.16-alpine AS build

# Установка зависимостей
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование файлов в рабочую директорию
COPY . .

# Получение зависимостей
RUN go mod download

# Сборка приложения
RUN go build -o /go/bin/bot ./cmd/bot

# Базовый образ для запуска приложения
FROM alpine:3.13

# Установка зависимостей
RUN apk add --no-cache ca-certificates

# Установка рабочей директории
WORKDIR /app

# Копирование бинарного файла из базового образа для сборки
COPY --from=build /go/bin/bot .

# Открытие порта для бота
EXPOSE 8080

# Команда для запуска приложения
CMD ["./bot"]

