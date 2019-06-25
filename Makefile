BINARY := bin/git-multi
GOBIN := $(GOPATH)/bin

build: $(BINARY)
$(BINARY): *.go
	go build -o $(BINARY) -v

install: build
	GOBIN=$(GOBIN) go install
