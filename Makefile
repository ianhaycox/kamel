.PHONY:

test-clean:
	go clean -testcache

test: generate
	go test -cover ./...

## cover-html: run coverage and show html report
cover-html:
	go test -cover ./... -coverprofile cover.out | grep -v "no test files"
	go tool cover -html=cover.out -o cover.html

## generate: runs go generate
generate:
	@command -v mockgen > /dev/null || \
		go install github.com/golang/mock/mockgen@v1.6.0
	go generate -v ./...

## clean-mock: removes all generated mocks
clean-mock:
	find . -iname '*_mock.go' -exec rm {} \;

## regenerate: clear and regenerate mocks
regenerate: clean-mock generate

## update: runs go mod vendor and tidy
update: mod tidy

## mod: runs go mod vendor
mod:
	go mod vendor -v

## tidy: runs go mod tidy
tidy:
	go mod tidy -v

lint:
	golangci-lint --version
	golangci-lint run -v ./...
