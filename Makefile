.PHONY: test lint targets/alexandrium

GINKGO := go run github.com/onsi/ginkgo/ginkgo

test: lint targets/alexandrium
	$(GINKGO) --race --randomizeAllSpecs -r .

lint:
	go vet ./...
	golangci-lint run ./...

targets/alexandrium:
	go build -o targets/alexandrium .

generate:
	go generate ./...