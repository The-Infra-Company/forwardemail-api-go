default: format check lint test

fmt:
	gofumpt -w .

check:
	go vet ./...

test:
	go test -v ./... -cover -count=1

vendor:
	go mod vendor

# Make sure you have installed golangci-lint CLI
# https://golangci-lint.run/usage/install/#local-installation
lint:
	golangci-lint run

.PHONY: vendor
