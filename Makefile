.PHONY: lint test-cover test-cover-svg test-cover-html install-deps
.SILENT:

lint:
	@#clear
	@$(CURDIR)/bin/golangci-lint run -c .golangci.yaml --path-prefix . --fix

test:
	@#clear
	@go test --cover --coverprofile=coverage.out $(TEST_COVER_EXCLUDE_DIR) -count=1

test-cover:
	@#clear
	@go test --coverprofile=coverage.out $(TEST_COVER_EXCLUDE_DIR) > /dev/null
	@go tool cover -func=coverage.out | grep total | grep -oE '[0-9]+(\.[0-9]+)?%'

test-cover-svg:
	@#clear
	@go test --coverprofile=coverage.out $(TEST_COVER_EXCLUDE_DIR) > /dev/null
	@$(CURDIR)/bin/go-cover-treemap -coverprofile coverage.out > coverage.svg
	@xdg-open ./coverage.svg

test-cover-html:
	@#clear
	@go test --coverprofile=coverage.out $(TEST_COVER_EXCLUDE_DIR) > /dev/null
	@go tool cover -html="coverage.out"

install-deps:
	@#clear
	@GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	@GOBIN=$(CURDIR)/bin go install github.com/nikolaydubina/go-cover-treemap@v1.3.0
	@GOBIN=$(CURDIR)/bin go install golang.org/x/tools/cmd/godoc@latest
	@go mod tidy

# ---------------

TEST_COVER_EXCLUDE_DIR := `go list ./... | grep -v -E 'cmd$$'`