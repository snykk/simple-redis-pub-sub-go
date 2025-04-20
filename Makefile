# Makefile for simple redis pub/sub go

# Services
SERVICES := order payment inventory notifier

# Directories
SRC_DIR := ./services
SHARED_DIR := ./shared

# Docker setup
DOCKER_COMPOSE := docker-compose -f docker-compose.yml
REDIS_CONTAINER_NAME := redis

# Default target to run when no target is specified
.DEFAULT_GOAL := help

# Target for starting Redis container
start-redis:
	$(DOCKER_COMPOSE) up -d redis
	@echo "Redis started."

# Target for stopping Redis container
stop-redis:
	$(DOCKER_COMPOSE) down
	@echo "Redis stopped."

# Run all services
# run-all: run-order run-payment run-inventory run-notifier

# Run a specific service 
run-%:
	@echo "Running $* service..."
	go run $(SRC_DIR)/$*/main.go
	@echo "$* service is running."

# Run all listeners
# run-listener: run-payment run-inventory run-notifier
# run-listener:
# 	@echo "Starting payment listener..."
# 	go run $(SRC_DIR)/payment/main.go &

# 	@echo "Starting inventory listener..."
# 	go run $(SRC_DIR)/inventory/main.go &

# 	@echo "Starting notifier listener..."
# 	go run $(SRC_DIR)/notifier/main.go &

# 	@echo "All listeners are running in the background."


# Run order service and publish a message
run-create-order:
	@echo "Creating order..."
	go run $(SRC_DIR)/order/main.go
	@echo "Order created and message published."

# Show all available targets
help:
	@echo "Makefile help"
	@echo
	@echo "Targets:"
	@echo "  run-<service>    Run a specific service (e.g., run-order)."
	@echo "  run-create-order Run the create order process."
	@echo "  start-redis      Start Redis using Docker."
	@echo "  stop-redis       Stop Redis."
	@echo "  help             Show this help message."