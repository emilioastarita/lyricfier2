.PHONY: clean run build release format releases release-linux release-windows release-darwin
VERSION = $$(git describe --abbrev=0 --tags)
VERSION_DATE = $$(git log -1 --pretty='%ad' --date=format:'%Y-%m-%d' $(VERSION))
COMMIT_REV = $$(git rev-list -n 1 $(VERSION))
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt
GOGENERATE=$(GOCMD) generate
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_FOLDER=$(realpath build/)
BINARY_NAME=$(BUILD_FOLDER)/lyricfier
STATICS_DIR=$(realpath lyricfier/)
GO_SOURCES:=$(shell find lyricfier/ -type f -name '*.go')
GO_SOURCES_INTERNAL:=$(shell find internal/ -type f -name '*.go')
STATIC_EMBEDED:=internal/lyricfier/static.go
STATIC_SOURCES:=$(shell find lyricfier/static/ -type f -name '*')

build: $(BINARY_NAME) $(STATIC_EMBEDED) .deps_updated

release-windows:  $(GO_SOURCES) $(GO_SOURCES_INTERNAL) $(STATIC_EMBEDED)
	cd lyricfier/ ; env GOOS=windows GOARCH=amd64 $(GOBUILD) -v  -ldflags -H=windowsgui -o $(BINARY_NAME)-windows-amd64.exe ; cd -

release-darwin:  $(GO_SOURCES) $(GO_SOURCES_INTERNAL) $(STATIC_EMBEDED)
	cd lyricfier/ ; env GOOS=darwin GOARCH=amd64 $(GOBUILD) -v   -o $(BINARY_NAME)-darwin-amd64 ; cd -

release-linux:  $(GO_SOURCES) $(GO_SOURCES_INTERNAL) $(STATIC_EMBEDED)
	cd lyricfier/ ; env GOOS=linux GOARCH=amd64 $(GOBUILD) -v   -o $(BINARY_NAME)-linux-amd64 ; cd -

releases: release-darwin release-linux release-windows

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

