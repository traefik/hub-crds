default: build

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build
build: lint test
	@$(CURDIR)/script/code-gen.sh
