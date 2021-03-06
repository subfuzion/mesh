PKG := github.com/subfuzion/mesh
CMD := $(PKG)/cmd
ASSETS := $(shell find assets -type f -print)

build: bin/mesh
rebuild: generate build
all: clean generate check build

bin/mesh: assets_vfsdata.go
	go build -o $@ "$(CMD)/$(@F)"

assets_vfsdata.go: $(ASSETS)
	go generate

.PHONY: generate
generate: assets_vfsdata.go

PHONY: dev
dev:
	go build -o bin/mesh -tags dev $(CMD)/mesh

.PHONY fmt:
	@gofmt -d -e -l -s -w .

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint:
	@which golint >/dev/null || go get -u golang.org/x/lint/golint
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

