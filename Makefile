.PHONY = test lint

GINKGO := go run github.com/onsi/ginkgo/ginkgo

test: lint
	$(GINKGO) --race --randomizeAllSpecs -r .

lint:
	go vet ./...
	golangci-lint run ./...