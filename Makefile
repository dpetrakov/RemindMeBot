# Переменные
APP_NAME = RemindMeBot
DOCKER_COMPOSE_FILE = docker-compose.yml

# Сборка Docker-образов
build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Запуск контейнеров
run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Остановка контейнеров
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Пересборка контейнеров и запуск приложения
restart: stop build run

# Очистка ненужных Docker-ресурсов
clean:
	docker system prune --force --volumes

# Тестирование кода
test:
	go test -v ./...

# Проверка кода линтером
lint:
	golint ./...

# Проверка кода на наличие ошибок
vet:
	go vet ./...

# Сборка и установка приложения на локальной машине
install:
	go install ./cmd/bot

# Отображение версии приложения
version:
	@echo "$(APP_NAME) version 1.0"

# Помощь по доступным командам
help:
	@echo "Usage:"
	@echo " make build - сборка Docker-образов"
	@echo " make run - запуск контейнеров"
	@echo " make stop - остановка контейнеров"
	@echo " make restart - перезапуск контейнеров"
	@echo " make clean - очистка ненужных Docker-ресурсов"
	@echo " make test - тестирование кода"
	@echo " make lint - проверка кода линтером"
	@echo " make vet - проверка кода на наличие ошибок"
	@echo " make install - сборка и установка приложения на локальной машине"
	@echo " make version - отображение версии приложения"
	@echo " make help - помощь по доступным командам"

