snapshot-single: clean
	goreleaser build --snapshot --single-target

snapshot: clean
	goreleaser build --snapshot

build: clean
	goreleaser build

release: clean
	goreleaser release --clean

generate:
	go generate ./...

lint:
	golangci-lint run --timeout 5m

test:
	go test -coverprofile coverage.out -race ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

clean:
	go clean
	rm -rf dist/
