VERSION=1.0.0
SOURCE_FILES=$(shell find . -type d -name vendor -prune -o -type d -path ./cmd -prune -o -type f -name '*.go' -print)
GO_LIST=$(shell go list ./... | grep -v /vendor/)
GOPATH=$(shell echo "$$GOPATH")

all: tree

dist: build-cross
	cd bin/linux/amd64 && tar zcvf realize_sample-linux-amd64-${VERSION}.tar.gz realize_sample-${VERSION}

tree: ${SOURCE_FILES} main.go
	go build -o bin/tree main.go

bundle:
	dep ensure

checkall: lint vet errcheck misspell

fmt:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 gofmt -w

fmtcheck:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 gofmt -l | xargs -r false

import:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 goimports -w

importcheck:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 goimports -l | xargs -r false

lint:
	golint -set_exit_status ${GO_LIST}

vet:
	go vet ${GO_LIST}

errcheck:
	errcheck -blank -ignoretests ${GO_LIST}

misspell:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 misspell -error

test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out -v ./...

viewcover:
	go tool cover -html=cover.out

clean:
	rm -rf bin/*
