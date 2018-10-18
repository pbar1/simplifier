GOCMD ?= go
BINARY_NAME = simplifier
PLATFORMS = linux darwin windows

build:
	$(GOCMD) build -o bin/$(BINARY_NAME)

startredis:
	docker-compose up -d redis

stopredis:
	docker-compose stop redis

.PHONY: build
