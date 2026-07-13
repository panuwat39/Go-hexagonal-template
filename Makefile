APP_NAME ?= app

.PHONY: run dev test test-race build fmt fmt-check vet tidy check vulncheck new-module init-project clean

run:
	go run ./cmd/api

dev:
	./scripts/dev.sh

test:
	go test ./...

test-race:
	go test -race -cover ./...

build:
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd/api

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Run make fmt"; gofmt -l .; exit 1)

vet:
	go vet ./...

tidy:
	go mod tidy

check:
	./scripts/check.sh

vulncheck:
	@command -v govulncheck >/dev/null 2>&1 || go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

new-module:
	@test -n "$(name)" || (echo "Usage: make new-module name=user"; exit 1)
	./scripts/new-module.sh $(name)

init-project:
	@test -n "$(module)" || (echo "Usage: make init-project module=github.com/you/app name=app"; exit 1)
	./scripts/init-project.sh $(module) $(name)

clean:
	rm -rf bin
