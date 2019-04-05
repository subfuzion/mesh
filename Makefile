PKG := github.com/subfuzion/meshdemo
CMD := $(PKG)/cmd
ASSETS := $(shell find assets -type f -print)

build: bin/meshdemo
rebuild: clean build
all: check clean build

bin/meshdemo: generate
	go build -o $@ "$(CMD)/$(@F)"

assets_vfsdata.go: $(ASSETS)
	go generate

generate: assets_vfsdata.go

PHONY: dev
dev:
	go build -o bin/meshdemo -tags dev $(CMD)/meshdemo

.PHONY fmt:
	@gofmt -d -e -l -s -w .

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint:
	@golint ./...

.PHONY: check
check: fmt vet lint

.PHONY: clean
clean:
	rm -rf bin assets_vfsdata.go

.PHONY: goget
goget:
	go get -u golang.org/x/lint/golint
	go get -u github.com/shurcooL/vfsgen

