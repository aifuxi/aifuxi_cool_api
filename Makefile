.PHONY: server build

server:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app	