GOPROXY ?= https://goproxy.cn,direct
GO_RUN = $(if $(GOPROXY),GOPROXY=$(GOPROXY),)
GOLANGCI_LINT_VERSION ?= v2.11.4

.PHONY: dev test lint fmt build frontend-test frontend-build fullstack docs-check docker-build ci-check

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

docs-check:
	node -e "const fs=require('fs');const path=require('path');const docs=[...walk(path.join(process.cwd(),'ai_docs')),path.join(process.cwd(),'README.md'),path.join(process.cwd(),'AGENTS.md'),path.join(process.cwd(),'CLAUDE.md')];const pattern=/\\[[^\\]]+\\]\\(([^)#]+)\\)/g;const missing=[];for(const doc of docs){const text=fs.readFileSync(doc,'utf8');for(const match of text.matchAll(pattern)){const rel=match[1];if(rel.includes('://')||rel.startsWith('#')) continue;const target=path.resolve(path.dirname(doc),rel);if(!fs.existsSync(target)) missing.push(`${doc}: ${rel}`);}}if(missing.length>0){console.error(missing.join('\\n'));process.exit(1);}function walk(dir){const entries=fs.readdirSync(dir,{withFileTypes:true});const files=[];for(const entry of entries){const fullPath=path.join(dir,entry.name);if(entry.isDirectory()){files.push(...walk(fullPath));continue;}if(entry.isFile()&&fullPath.endsWith('.md')){files.push(fullPath);}}return files;}"

docker-build:
	docker build -t go-template:local .

ci-check: fullstack docs-check docker-build
