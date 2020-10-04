.PHONY: test lint targets/alexandrium

GINKGO := go run github.com/onsi/ginkgo/ginkgo

test: lint targets/alexandrium
	$(GINKGO) --race --randomizeAllSpecs -r .

lint: vet
# In CI, this is run via the official action
ifndef SKIP_LINT
	golangci-lint run ./...
endif

vet:
	go vet ./...

targets/alexandrium:
	go build -o targets/alexandrium .

generate:
	go generate ./...