.DEFAULT_GOAL := all

all: vet lint test build
.PHONY: all

build: vet
	go build -o ./cmd/diabolus/diabolus ./cmd/diabolus/main.go
.PHONY: build

clean:
	@rm -f ./cmd/diabolus/diabolus
.PHONY: clean

fmt:
	goimports -l -w .
.PHONY: fmt

lint: fmt
	staticcheck ./...
.PHONY: lint

test:
	go test ./...
.PHONY: test

vet: fmt
	go vet ./...
	shadow ./...
.PHONY: vet
