BIN_NAME := mvnc

out/$(BIN_NAME): go.mod go.sum $(wildcard *.go) $(wildcard */*.go)
	go build -o out/$(BIN_NAME)

.PHONY: check
check:
	go vet  ./...
	go test ./...

.PHONY: clean
clean:
	go clean ./...
	-rm -rf out
