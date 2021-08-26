PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin/
BINARY_MAC_NAME=main
GOMAIN=$(GOBASE)/cmd/ova-hobby-api/main.go

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: run, test, check, build, all, help

## run: Run a main.go
run: 
	go run $(GOMAIN)

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

## help: Show help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
