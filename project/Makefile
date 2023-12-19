FRONTEND_BINARY=front-endApp
LOCAL_FRONTEND_BINARY=webApp
BROKER_BINARY=brokerApp
QUOTE_SERVICE_BINARY=quoteServiceApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_frontend build_broker build_quoteservice
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_frontend: builds the fontend binary as a linux executable
build_frontend:
	@echo "Building front-end binary..."
	cd ../frontend && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONTEND_BINARY} ./cmd/web
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_quoteservice: builds the quote-service binary as a linux executable
build_quoteservice:
	@echo "Building quote-service binary..."
	cd ../quote-service && env GOOS=linux CGO_ENABLED=0 go build -o ${QUOTE_SERVICE_BINARY} ./cmd/api
	@echo "Done!"

## build_local_front: builds the frone end binary
build_local_front:
	@echo "Building front end binary..."
	cd ../frontend && env CGO_ENABLED=0 go build -o ${LOCAL_FRONTEND_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_local_front
	@echo "Starting front end"
	cd ../frontend && ./${LOCAL_FRONTEND_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${LOCAL_FRONTEND_BINARY}"
	@echo "Stopped front end!"