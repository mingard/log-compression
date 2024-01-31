GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
COVERPROFILE=cover.out
COVERALLS_TOKEN=
GOTEST=$(GOCMD) test -cover -coverprofile $(COVERPROFILE)
GOGET=$(GOCMD) get

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

SRC ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS = $(shell go list ./... | grep -v /vendor)

.DEFAULT_GOAL: lint
.PHONY: install add update proto hooks fmt test coverage_badge coverage_report lint

%:
	@true

install:
	dep ensure -v

add:
	echo $(MAKECMDGOALS) | cut -d ' ' -f2- | xargs dep ensure -add

update:
	echo $(MAKECMDGOALS) | cut -d ' ' -f2- | xargs dep ensure -update

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
		go get -u -v github.com/golang/protobuf/protoc-gen-go
		for file in $$(git ls-files '*.proto'); do \
			protoc -I $$(dirname $$file) --go_out=plugins=grpc:$$(dirname $$file) $$file; \
		done

hooks: .git/hooks/pre-commit

.git/hooks/pre-commit: scripts/hooks/pre-commit.sh
	@cp -f scripts/hooks/pre-commit.sh .git/hooks/pre-commit

fmt:
	goimports -l -w $(SRC)

test:
	$(GOTEST) $(PKGS)

coverage_badge:
	goveralls -coverprofile=$(COVERPROFILE) -repotoken $(COVERALLS_TOKEN)

coverage_report:
	gocov convert $(COVERPROFILE) | gocov annotate -

lint: $(GOMETALINTER)
	$(GOMETALINTER) ./... --vendor --skip api --deadline "60s" --enable=goimports

$(GOMETALINTER):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install 1>/dev/null
