GO ?= godep go
COVERAGEDIR = ./coverage
all: build test cover

godep:
	go get github.com/tools/godep

godep-save:
	godep save ./...

all: build test

build:
	$(GO) build -v ./...

fmt:
	$(GO) fmt ./...

test:
	if [ ! -d $(COVERAGEDIR) ]; then mkdir $(COVERAGEDIR); fi
	$(GO) test -v ./loco -race -cover -coverprofile=$(COVERAGEDIR)/loco.coverprofile

cover:
	$(GO) tool cover -html=$(COVERAGEDIR)/loco.coverprofile -o $(COVERAGEDIR)/loco.html

bench:
	$(GO) test ./... -cpu 2 -bench .
