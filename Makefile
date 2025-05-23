.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: test
test:
	go test -v -race -coverprofile="coverage.txt" -covermode=atomic ./...

.PHONY: release
release:
	goreleaser build --snapshot
