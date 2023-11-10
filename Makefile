default: build

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: lint
	@$(CURDIR)/script/code-gen.sh
