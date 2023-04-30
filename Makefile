.PHONY: build
build: deps lint compile test install

.PHONY: deps
deps:
	@ go mod tidy --compat=1.20

.PHONY: compile
compile:
	@ go build -o bin/lsgo cmd/lsgo/main.go

.PHONY: install
install:
	go install ./cmd/lsgo

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml --fix

.PHONY: run
run:
	go run ./cmd/lsgo/main.go