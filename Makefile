# Variables
APP_NAME := knovault

templ-generate:
	TEMPL_EXPERIMENT=rawgo templ generate

dev: templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

build: templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

docker-dev-build:
	docker build --network=host --no-cache -t knovault-dev -f Dockerfile_dev .

# Use shell command to get absolute path
PWD := $(shell pwd)

docker-dev-run:
	docker run -it --rm \
	-v "${PWD}:/app" \
	-p 1323:1323 \
	-w /app \
	--add-host=proxy.golang.org:172.217.22.113 \
	knovault-dev

.PHONY: clean-assets compile-plugin compile-theme templ-generate ensure-dirs dev build docker-dev-build docker-dev-run