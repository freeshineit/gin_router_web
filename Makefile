#一定不能用4个空格代替tab

GOCMD=go
GOBUILD=$(GOCMD) build -v -ldflags '-w -s'
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
GOGET=$(GOCMD) get -u -v

OS := $(shell uname -s | awk '{print tolower($$0)}')

BINARY = ./bin/app


LDFLAGS = -ldflags="$$(govvv -flags)"

all: clean build

test:
	$(GOTEST)

lint:
	golint

build:
	env CGO_ENABLED=0 GOOS=$(OS) GIN_MODE=release  $(GOBUILD) -o $(BINARY) ./main.go

clean:
	$(GOCLEAN)
	@rm -f $(BINARY)

deploy:
	docker build -t gin_router_web:latest .
# docker network create my_net
# docker-compose up -d
	docker compose up -d
