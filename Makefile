default: build lint test

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

clean:
	find pkg/apis/hub/ -name '*.yaml' -delete

.PHONY: build
build: clean
	@$(CURDIR)/script/code-gen-docker.sh
