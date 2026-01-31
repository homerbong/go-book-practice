.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

# As part of exercise 2
clean:
	go clean

# Expanding on exercise 2
run: build
	./hello_world