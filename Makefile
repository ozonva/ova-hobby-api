include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin/
BINARY_MAC_NAME=main
GOMAIN=$(GOBASE)/cmd/ova-hobby-api/main.go

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: run, test, check, build, all, help, up, down

## migrate: Apply migrations to the DB
migrate:
	goose -dir ./migrations/ postgres "user=$(POSTGRES_USER) dbname=$(POSTGRES_DB) password=$(POSTGRES_PASSWORD) sslmode=disable" up


## up: Run docker containers
up:
	docker compose up --build -d

## down: Stop and remove docker containers
down:
	docker compose down --remove-orphans

## run: Run a main.go
run:
	source ./.env && go run $(GOMAIN)

## test: Test the project
test:
	go test -race ./...

## check: Run formatting and code checks
check:
	go fmt ./...
	staticcheck ./...

## build: Build a binary
build:
	go build -o $(GOBIN)$(BINARY_MAC_NAME) $(GOMAIN)

## generate: Generate everything
generate:
	go generate ./...

## all: Code checks and run
all: check run

.PHONY: deps
deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init github.com/ozonva/ova-hobby-api
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger


.PHONY: proto
proto:
	protoc -I api/ \
			--go_out=pkg/ --go_opt=paths=import \
			--go-grpc_out=pkg/ --go-grpc_opt=paths=import \
			api/hobby.proto

## help: Show help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo