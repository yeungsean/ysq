
include Makefile.mk

test:
	@$(GO) test -gcflags=all=-l -timeout 30s -covermode=count ./...
