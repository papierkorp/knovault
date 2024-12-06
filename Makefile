# Variables
APP_NAME := knovault
ASSET_DIR := ./internal/assetManager

clean-assets:
	@echo "Cleaning up asset artifacts..."
	@find $(ASSET_DIR) -name "*.so" -type f -delete

compile-plugin:
	@if [ -z "$(PLUGIN)" ]; then \
		echo "Usage: make compile-plugin PLUGIN=PluginName"; \
		exit 1; \
	fi
	@echo "Compiling plugin $(PLUGIN)..."
	@cd $(ASSET_DIR)/plugins/$(PLUGIN) && \
	go build -buildmode=plugin -o $(PLUGIN).so main.go

compile-theme:
	@if [ -z "$(THEME)" ]; then \
		echo "Usage: make compile-theme THEME=ThemeName"; \
		exit 1; \
	fi
	@echo "Compiling theme $(THEME)..."
	@cd $(ASSET_DIR)/themes/$(THEME) && \
	go build -buildmode=plugin -o $(THEME).so main.go

templ-generate:
	TEMPL_EXPERIMENT=rawgo templ generate

ensure-dirs:
	@mkdir -p $(ASSET_DIR)/plugins $(ASSET_DIR)/themes

dev: ensure-dirs templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

build: ensure-dirs templ-generate
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