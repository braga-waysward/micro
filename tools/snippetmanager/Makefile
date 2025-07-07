# Makefile for Micro Snippet Manager
# Provides targets for building, installing, and testing the tool

.PHONY: all build install uninstall clean test

# Build settings
BINARY_NAME := snippetmanager
INSTALL_DIR := /usr/local/bin
CONFIG_DIR := $(HOME)/.config/micro

all: build

# Build the snippet manager binary
build:
	@echo "Building snippet manager..."
	go build -o $(BINARY_NAME) snippetmanager.go
	@echo "Build complete."

# Install the binary and ensure config directory exists
install: build
	@echo "Installing snippet manager..."
	mkdir -p $(CONFIG_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)
	@echo "Installation complete. Snippets will be stored in $(CONFIG_DIR)/snippets.json"

# Remove the installed binary
uninstall:
	@echo "Uninstalling snippet manager..."
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete. Note: Snippet data in $(CONFIG_DIR) is preserved."

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	@echo "Clean complete."

# Run basic tests (this would be expanded in a real project)
test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Tests complete."