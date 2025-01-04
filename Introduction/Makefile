.DEFAULT_GOAL := run

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

run: build
	./bin/hello
	@echo "run done"

clean: build
	go clean -i
	rm -rf bin
	@echo "clean done"
