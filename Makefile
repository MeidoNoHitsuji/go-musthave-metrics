.PHONY: build
build:
	cd ./cmd/server && \
	go build -o server *.go
	cd ./cmd/agent && \
	go build -o agent *.go