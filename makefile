build: format
	docker-compose build --no-cache

build_linux: format
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o build/knowledge -v ./cmd/knowledge

down:
	docker-compose down

format:
	go fmt ./...

start:
	docker-compose start

stop:
	docker-compose stop

up:
	docker-compose up -d

.PHONY: \
	build \
	build_linux \
	down \
	format \
	start \
	stop \
	up