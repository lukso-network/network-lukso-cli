DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match="v*" 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BRANCH  ?= $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)

test: build
	./lukso --help
	./lukso version
	./lukso cluster sync

build:
	@CGO_ENABLED=0 go build \
		-ldflags '-X main.Version=$(VERSION) -X main.BuildDate=$(DATE) -X main.GitCommit=$(COMMIT) -X main.GitBranch=$(BRANCH)' \
		-o lukso ./cmd/lukso/*.go