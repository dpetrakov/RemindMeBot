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
RUN go build -o /go/bin/reminder ./cmd/reminder

# Базовый образ для запуска приложения
FROM alpine:3.13

# Установка зависимостей
RUN apk add --no-cache ca-certificates

# Установка рабочей директории
WORKDIR /app

# Копирование бинарного файла из базового образа для сборки
COPY --from=build /go/bin/reminder .

# Команда для запуска приложения
CMD ["./reminder"]

