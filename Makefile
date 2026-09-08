GOPROXY ?= https://goproxy.cn,direct
GO_RUN = $(if $(GOPROXY),GOPROXY=$(GOPROXY),)
GOLANGCI_LINT_VERSION ?= v2.11.4

.PHONY: dev test lint fmt build frontend-test frontend-build fullstack docker-build ci-check

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

frontend-test:
	pnpm --prefix frontend test --run

frontend-build:
	pnpm --prefix frontend build

fullstack: test build lint frontend-test frontend-build

docker-build:
	docker build -t go-template:local .

ci-check: fullstack docker-build
