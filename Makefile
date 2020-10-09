.PHONY: test lint targets/alexandrium

GINKGO := go run github.com/onsi/ginkgo/ginkgo

DATABASE_URL ?= "postgresql://postgres:foo@127.0.0.1:5432/postgres"
test: lint targets/alexandrium
	@$(GINKGO) --race --randomizeAllSpecs -r store/memory/ router/ acceptance/
ifndef SKIP_RUN_POSTGRES
	@docker run --rm --name alexandrium-pg-test -p 5432:5432 -e POSTGRES_PASSWORD=foo -d postgres
endif
	@DATABASE_URL=$(DATABASE_URL) $(GINKGO) --race --randomizeAllSpecs -r store/postgres/
ifndef SKIP_RUN_POSTGRES
	@docker stop alexandrium-pg-test
endif

lint: vet
# In CI, this is run via the official action
ifndef SKIP_LINT
	@golangci-lint run ./...
endif

vet:
	@go vet ./...

targets/alexandrium:
	@go build -o targets/alexandrium .

generate:
	@go generate ./...