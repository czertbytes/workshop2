# Project
PROJECT?=articles
ORGANIZATION?=czertbytes
REPOSITORY?=github.com

# Build
GO_BUILD_CMD?=go build -o bin/articles .

# Project go deps
GO_DEPS?=golang.org/x/net/html

all: build-cli

## Dependencies and vendor
deps:
	for dep in $(GO_DEPS); do \
		go get -u $$dep ; \
	done

clean-vendor:
	rm -rf vendor/

vendor: clean-vendor
	for dep in $(GO_DEPS); do \
		mkdir -p vendor/$$dep ; \
		cp -R $(GOPATH)src/$$dep/* vendor/$$dep/ ; \
		rm -rf vendor/$$dep/.git ; \
	done

clean:
	rm -rf bin/

build-cli: clean
	mkdir -p bin
	$(GO_BUILD_CMD)

.PHONY: all