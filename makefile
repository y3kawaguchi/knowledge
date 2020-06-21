build:
	go build -o build/knowledge -v ./cmd/knowledge/

build_linux:
	GOOS=linux GOARCH=amd64 go build -o build/knowledge -v ./cmd/knowledge

format:
	go fmt ./...

run:
	build/knowledge

.PHONY: \
	build \
	build_linux \
	format \
	run