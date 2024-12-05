# Variables
APP_NAME := knovault
THEMES_DIR := ./internal/themes
PLUGINS_DIR := ./internal/plugins

# Common clean function for both plugins and themes
define clean_artifacts
	@echo "Cleaning up $(1) artifacts..."
	@find $(2)/core -name "internal" -type d -exec rm -rf {} +
	@find $(2)/core -name "*.so" -type f -delete
endef

# Common build function for both plugins and themes
define build_components
	@echo "Building $(1)..."
	@for item in $(2)/core/*; do \
		if [ -d "$$item" ]; then \
			name=$$(basename $$item); \
			echo "Building $$name..."; \
			cd $$item && go build -buildmode=plugin -o $$name.so main.go && cd -; \
		fi \
	done
endef

.PHONY: clean-plugins clean-themes templ-generate build-themes build-plugins ensure-dirs dev build docker-build-base docker-build-dev docker-build-prod docker-run-dev docker-run-prod docker-stop

clean-plugins:
	$(call clean_artifacts,plugin,$(PLUGINS_DIR))

clean-themes:
	$(call clean_artifacts,theme,$(THEMES_DIR))

templ-generate:
	TEMPL_EXPERIMENT=rawgo templ generate

build-themes: clean-themes
	$(call build_components,theme plugins,$(THEMES_DIR))

build-plugins: clean-plugins
	$(call build_components,plugins,$(PLUGINS_DIR))

ensure-dirs:
	@mkdir -p $(THEMES_DIR)/core $(THEMES_DIR)/common $(PLUGINS_DIR)/core $(PLUGINS_DIR)/common

dev: ensure-dirs templ-generate build-themes build-plugins
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

build: ensure-dirs templ-generate build-themes build-plugins
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

docker-build-base:
	@if [ -z "$$(docker images -q knovault_base:latest 2> /dev/null)" ]; then \
		echo "Building knovault_base image..."; \
		docker build -t knovault_base:latest -f Dockerfile_base .; \
	else \
		echo "knovault_base image already exists. Skipping build."; \
	fi

docker-build-%: docker-build-base
	docker build -t knovault:$* -f Dockerfile.$* .

docker-run-%:
	docker run -d --name knovault-$* -p 1323:1323 \
		$(if $(findstring dev,$*),-v $(PWD):/app,-v $(PWD)/data:/app/data) \
		knovault:$*

docker-stop:
	-docker stop knovault-dev knovault-prod
	-docker rm knovault-dev knovault-prod