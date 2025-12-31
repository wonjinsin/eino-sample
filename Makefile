PACKAGE = github.com/wonjinsin/eino-sample
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BIN = $(BASE_PATH)/bin
BINARY_NAME = bin/server
MAIN = $(BASE_PATH)/cmd/server/main.go
GOLINT = $(BIN)/golint
GOBIN = $(shell go env GOPATH)/bin
MOCK = $(GOBIN)/mockgen
PKG_LIST = $(shell cd $(BASE_PATH) && cat pkg.list)


ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif

.PHONY: tool
tool:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install go.uber.org/mock/mockgen@latest

.PHONY: build
build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

.PHONY: vet
vet:
	go vet

.PHONY: fmt
fmt:
	go fmt

.PHONY: lint
lint:
	$(GOBIN)/golangci-lint run

.PHONY: test
test: build-mocks
	go test -v -cover ./...

.PHONY: test-all
test-all: test vet fmt lint

.PHONY: build-mocks
build-mocks:
	$(MOCK) -source=internal/usecase/service.go -destination=mock/mock_service.go -package=mock
	$(MOCK) -source=internal/repository/repository.go -destination=mock/mock_repository.go -package=mock

.PHONY: init
init: 
	go mod init $(PACKAGE)

.PHONY: tidy
tidy: 
	go mod tidy

.PHONY: vendor
vendor: build-mocks
	go mod vendor

.PHONY: start
start: build 
	@$(BINARY_NAME)

.PHONY: all
all: tool init tidy vendor build

.PHONY: clean
clean:; $(info cleaning…) @ 
	@rm -rf vendor mock bin
	@rm -rf go.mod go.sum pkg.list
