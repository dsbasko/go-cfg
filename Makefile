.PHONY: lint install-deps
.SILENT:

lint:
	@clear
	@$(CURDIR)/bin/golangci-lint run -c .golangci.yaml --path-prefix .

install-deps:
	@clear
	@GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	@go mod tidy

