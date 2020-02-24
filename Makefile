.PHONY: clean run build release format
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt
GOGENERATE=$(GOCMD) generate
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=$(realpath build/lyricfier)
STATICS_DIR=$(realpath lyricfier/)
GO_SOURCES:=$(shell find lyricfier/ -type f -name '*.go')
GO_SOURCES_INTERNAL:=$(shell find internal/ -type f -name '*.go')
STATIC_EMBEDED:=internal/lyricfier/static.go
STATIC_SOURCES:=$(shell find lyricfier/static/ -type f -name '*')

build: $(BINARY_NAME) $(STATIC_EMBEDED) .deps_updated

build-windows:  $(GO_SOURCES) $(GO_SOURCES_INTERNAL) $(STATIC_EMBEDED)
	env GOOS=windows GOARCH=amd64 $(GOBUILD) -v -o $(BINARY_NAME)-amd64.exe $(GO_SOURCES)

$(BINARY_NAME): $(GO_SOURCES) $(GO_SOURCES_INTERNAL) $(STATIC_EMBEDED)
	$(GOBUILD) -o $(BINARY_NAME) -v $(GO_SOURCES)

format:
	$(GOFMT) $(GO_SOURCES)
	$(GOFMT) internal/search/*.go
	$(GOFMT) internal/lyricfier/*.go

$(STATIC_EMBEDED): $(STATIC_SOURCES)
	$(GOGENERATE) -v $(GO_SOURCES)

.deps_updated: $(GO_SOURCES)
	$(GOCMD) mod tidy
	touch .deps_updated

run: build
	cd $(STATICS_DIR) ; env LOCAL_ASSETS=true $(BINARY_NAME) ; cd -

clean:
	rm $(BINARY_NAME)

