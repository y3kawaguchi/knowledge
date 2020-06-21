build:
	go build -o build/knowledge -v ./cmd/knowledge/

format:
	go fmt ./...

run:
	build/knowledge

.PHONY: \
	build \
	format \
	run