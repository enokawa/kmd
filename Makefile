BINARY_NAME = kmd

export GO111MODULE=on

run:
	go run main.go
build:
	go build -o ./dist/$(BINARY_NAME)
