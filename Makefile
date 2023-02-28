.PHONY: all dev pre-dev clean pre-commit install-dev test build fmt

APP_NAME = go-tiger
BUILD_DIR = $(PWD)/build

all: dev

dev:
	mkdir -p build
	air

pre-dev:
	go mod tidy
	make build
	docker cp ./build/go-tiger joj2-go-tiger-1:.
	docker cp ./build/runner joj2-go-tiger-1:.
	docker cp ./build/go-tiger joj2-go-tiger-2:.
	docker cp ./build/runner joj2-go-tiger-2:.
	docker restart joj2-go-tiger-1
	docker restart joj2-go-tiger-2
	docker logs --since 0s -f joj2-go-tiger-1

clean:
	rm -rf ./build

pre-commit:
	pre-commit run --all-files

install-dev:
	go install github.com/cosmtrek/air@latest
	python3 -m pip install -U pre-commit
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	pre-commit install

test: clean
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

BUILT_AT := $(shell date +'%F %T %z')
GO_VERSION := $(shell go version)
GIT_AUTHOR := $(shell git show -s --format='format:%aN <%ae>' HEAD)
GIT_COMMIT := $(shell git log --pretty=format:"%H" -1)
VERSION := dev
FLAGS := "-s -w \
-X 'github.com/joint-online-judge/go-tiger/pkg/configs.BuiltAt=$(BUILT_AT)' \
-X 'github.com/joint-online-judge/go-tiger/pkg/configs.GoVersion=$(GO_VERSION)' \
-X 'github.com/joint-online-judge/go-tiger/pkg/configs.GitAuthor=$(GIT_AUTHOR)' \
-X 'github.com/joint-online-judge/go-tiger/pkg/configs.GitCommit=$(GIT_COMMIT)' \
-X 'github.com/joint-online-judge/go-tiger/pkg/configs.Version=$(VERSION)'"
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags=$(FLAGS) -o $(BUILD_DIR)/$(APP_NAME) .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags=$(FLAGS) -o $(BUILD_DIR)/runner ./pkg/runner

# run it manually as it is too time-consuming
fmt:
	golines -m 80 -w .
