GOPROXY ?= https://goproxy.cn,direct
GO_RUN = $(if $(GOPROXY),GOPROXY=$(GOPROXY),)
GOLANGCI_LINT_VERSION ?= v2.8.0

.PHONY: dev test lint fmt build

dev:
	$(GO_RUN) go run github.com/air-verse/air@latest

test:
	$(GO_RUN) go test ./...

lint:
	$(GO_RUN) go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run

fmt:
	$(GO_RUN) gofmt -w cmd internal

build:
	$(GO_RUN) go build ./...
