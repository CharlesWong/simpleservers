.PHONY: build test

default: build

build:  
	GOARCH=amd64 GOOS=linux go build -v -o bin/transproxy-amd64-linux transproxy/transproxy.go
	GOARCH=amd64 GOOS=linux go build -v -o bin/fileserver-amd64-linux fileserver/fileserver.go

test:
	go test ./...
