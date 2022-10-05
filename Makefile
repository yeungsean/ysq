
include Makefile.mk

test:
	@$(GO) test -gcflags=all=-l -timeout 30s -coverprofile=coverage.out ./...

mod:
	@$(GO) mod tidy
	git diff --exit-code go.mod
