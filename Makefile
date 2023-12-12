default: build lint test

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build
build:
	@$(CURDIR)/script/code-gen-docker.sh
