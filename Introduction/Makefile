.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt .
	@echo "fmt done"

vet: fmt
	go vet .
	@echo "vet done"

build: vet
	go build -o bin/hello .
	@echo "build done"
