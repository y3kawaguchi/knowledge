build:
	go build -o build/knowledge -v ./cmd/knowledge/

run:
	build/knowledge

.PHONY: \
	build \
	run