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

.PHONY: test1
test1:
	./metricstest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/server/server

.PHONY: test2
test2:
	./metricstest -test.v -test.run=^TestIteration2[AB]*$$ -source-path=. -agent-binary-path=cmd/agent/agent

.PHONY: test3
test3:
	./metricstest -test.v -test.run=^TestIteration3[AB]*$$ -source-path=. -agent-binary-path=cmd/agent/agent -binary-path=cmd/server/server
