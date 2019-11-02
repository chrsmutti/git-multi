NAME := git-multi
BINARY := dist/$(NAME)
GOBIN := $(GOPATH)/bin

build: $(BINARY)
$(BINARY): *.go
	go build -o $(BINARY) -v

install: build
	GOBIN=$(GOBIN) go install

clean:
	@rm -rf ./dist
