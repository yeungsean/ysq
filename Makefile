
include Makefile.mk

test:
	@$(GO) test -gcflags=all=-l -timeout=60s -coverprofile=coverage.txt -covermode=atomic -race -parallel 2 ./...

lint:
	@$(GOLANGCI-LINT) run ./...

mod:
	@$(GO) mod tidy
	git diff --exit-code go.mod

.PHONY: test, lint mod
