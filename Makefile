.PHONY: clean run build release
GOCMD=go
GOBUILD=$(GOCMD) build
GOGENERATE=$(GOCMD) generate
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=build/lyricfier
GO_SOURCES:=$(shell find lyricfier/ -type f -name '*.go')
STATIC_EMBEDED:=internal/lyricfier/static.go
STATIC_SOURCES:=$(shell find lyricfier/static/ -type f -name '*')

build: $(BINARY_NAME) $(STATIC_EMBEDED) .deps_updated

$(BINARY_NAME): $(GO_SOURCES)
	$(GOBUILD) -o $(BINARY_NAME) -v $(GO_SOURCES)

$(STATIC_EMBEDED): $(STATIC_SOURCES)
	$(GOGENERATE) -v $(GO_SOURCES)

.deps_updated: $(GO_SOURCES)
	$(GOCMD) mod tidy
	touch .deps_updated

run: build
	$(BINARY_NAME)

clean:
	rm $(BINARY_NAME)

