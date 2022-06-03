.PHONY: build run test

build: 
	go build main.go

run: build
	./main

test:
	go test ./...