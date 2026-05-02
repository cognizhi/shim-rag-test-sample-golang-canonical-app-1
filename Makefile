APP_NAME := canonical-calc

build:
	go build -o $(APP_NAME) ./cmd/canonicalcalc

test:
	go test ./...

run:
	go run ./cmd/canonicalcalc --help

clean:
	rm -f $(APP_NAME) coverage.out

.PHONY: build test run clean
