BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
DATE = $(shell git show --pretty=format:%cs --no-patch)
HASH = $(shell git show --pretty=format:%h --no-patch)
TAG = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%(refname:short)' --count=1)
GIT_INFO = $(BRANCH) - $(DATE) - $(HASH)

#env
DEV = dev
QA = qa
STAGE = stage
PROD = prod
#color
RED = \x1b[31;49;3;1m
GREEN = \x1b[32;49;3;1m
YELLOW = \x1b[33;49;3;1m
BLUE = \x1b[34;49;3;1m
MAGENTA = \x1b[35;49;3;1m
CYAN = \x1b[36;49;3;1m
WHITE = \x1b[49;3;1m

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

build-local: swag ## Build the local binary file
	@echo "${GREEN}build local binary file"
	@go build -tags local -v -a -ldflags="-X 'ggurugi/pkg/cli.GitInfo=$(GIT_INFO)' -X 'ggurugi/pkg/cli.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(DEV)'" -o bin/ggurugi main.go
.PHONY: build-local

build-dev: swag ## Build the dev binary file
	@echo "${GREEN}build dev binary file"
	@go build -tags dev -v -a -ldflags="-X 'ggurugi/pkg/cli.GitInfo=$(GIT_INFO)' -X 'ggurugi/pkg/cli.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(DEV)'" -o bin/ggurugi main.go
.PHONY: build-dev

build-qa: swag ## Build the qa binary file
	@echo "${GREEN}build qa binary file"
	@go build -tags qa -v -a -ldflags="-X 'ggurugi/pkg/cli.GitInfo=$(GIT_INFO)' -X 'ggurugi/pkg/cli.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(QA)'" -o bin/ggurugi main.go
.PHONY: build-qa

build-prod: swag ## Build the prod binary file
	@echo "${GREEN}build dev binary file"
	@go build -tags prod -a -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(PROD)'" -o bin/ggurugi main.go
.PHONY: build-prod

swag: ### swag init
	@echo "${RED}swag init"
	@swag init
.PHONY: swag

mock: ### run mockery
	@mockery --all --keeptree --dir $(PWD)/internal --output $(PWD)/mocks --disable-version-string
.PHONY: mock

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

run-dev: ### run the dev
	@echo "${YELLOW}run the swag binary file"
	go run -tags dev -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(DEV)'" main.go
.PHONY: run-dev

run-qa: ## run the qa
	@echo "${YELLOW}run the swag binary file"
	go run -tags qa -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(QA)'" main.go
.PHONY: run-qa

run-prod: ## run the prod
	@echo "${YELLOW}run the swag binary file"
	go run -tags prod -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(PROD)'" main.go
.PHONY: run-prod