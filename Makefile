SERVER_PORT:=$(shell ./get_free_port.sh)
ADDRESS:="localhost:${SERVER_PORT}"
TEMP_FILE:=./tmp

.PHONY: build
build:
	cd ./cmd/server && \
	go build -o server *.go
	cd ./cmd/agent && \
	go build -o agent *.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: test-all
test-all:
	make test1
	make test2
	make test3
	make test4
	make test5

.PHONY: test1
test1:
	./metricstest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/server/server

.PHONY: test2
test2:
	./metricstest -test.v -test.run=^TestIteration2[AB]*$$ -source-path=. -agent-binary-path=cmd/agent/agent

.PHONY: test3
test3:
	./metricstest -test.v -test.run=^TestIteration3[AB]*$$ -source-path=. -agent-binary-path=cmd/agent/agent -binary-path=cmd/server/server

.PHONY: test4
test4:
	./metricstest -test.v -test.run=^TestIteration4$$ -agent-binary-path=cmd/agent/agent -binary-path=cmd/server/server -server-port=$(SERVER_PORT) -source-path=.

.PHONY: test5
test5:
	./metricstest -test.v -test.run=^TestIteration5$$ -agent-binary-path=cmd/agent/agent -binary-path=cmd/server/server -server-port=$(SERVER_PORT) -source-path=.
