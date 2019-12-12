# Make is verbose in Linux. Make it silent.
# MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-$(MAKE) go-clean

run-reload: go-get run-service


run-service:
	@echo "  >  Running server..."
	reflex -r '\.go' -s -- /bin/sh -c 'SERVER_ADDR=:8080 SERVER_RETRY_PERIOD=10s go run ./server/server.go'

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get ./...

go-install:
	@go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@go clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo