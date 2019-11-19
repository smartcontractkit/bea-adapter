.DEFAULT_GOAL := build
.PHONY: build install docker

gomod:
    export GO111MODULE=on

build: gomod
	@go build -o cl-bea

install: gomod
	@go install

docker:
	@docker build . -t bea-adapter
