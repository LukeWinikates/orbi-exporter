.PHONY: test run

build/orbi-exporter-linux-amd64: $(shell find . -iname "*.go")
	mkdir -p build/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build  -o $@ cmd/main.go

test:
	go test -v ./...
	go vet ./...
	golangci-lint run

run:
	go run cmd/main.go
