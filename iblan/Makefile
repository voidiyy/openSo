# Назва вашого проекту
APP_NAME = iblan

# Визначення змінних
GO_FILES = $(wildcard cmd/**/*.go) $(wildcard internal/**/*.go)
UI_FILES = $(wildcard ui/**/*.html) $(wildcard ui/**/*.templ)
GO_MOD = go.mod
GO_SUM = go.sum

# Зовнішні команди
GO = go
GOINSTALL = $(GO) install
GOBUILD = $(GO) build
GORUN = $(GO) run
GOTEST = $(GO) test
GOCLEAN = $(GO) clean

# Цілі
.PHONY: all clean build run test install clean-mod

all: build

# Встановлення залежностей
install:
	$(GO) mod tidy

# Компіляція проекту
build: $(GO_FILES) $(UI_FILES)
	$(GOBUILD) -o $(APP_NAME) main.go

run: check-docker

	templ generate
	$(GORUN) main.go

check-docker:
	@if ! sudo docker ps | grep -q "iblanPostgres"; then \
		echo "Docker container 'iblanPostgres' is not running. Starting it now..."; \
		sudo docker start iblanPostgres; \
	fi

# Запуск тестів
test:
	$(GOTEST) ./...

# Очищення згенерованих файлів
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Видалення кешу модулів Go
clean-mod:
	$(GOCLEAN) -modcache

stop:
	@if sudo docker ps | grep -q "iblanPostgres"; then \
	sudo docker stop iblanPostgres; \
	echo "good job for today!"; \
	fi