BINARY_NAME=aoc2023

.PHONY: build
build:
	go build -o ${BINARY_NAME} ./cmd

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}

.PHONY: test
test:
	go test ./pkg/*

.PHONY: fmt
fmt:
	@gofumpt -l -w -e .
