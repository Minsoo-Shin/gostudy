BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
DATE = $(shell git show --pretty=format:%cs --no-patch)
HASH = $(shell git show --pretty=format:%h --no-patch)
TAG = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%(refname:short)' --count=1)
GIT_INFO = $(BRANCH) - $(DATE) - $(HASH)

#env
DEV = dev
PROD = prod
#color
RED = \x1b[31;49;3;1m
GREEN = \x1b[32;49;3;1m
YELLOW = \x1b[33;49;3;1m
BLUE = \x1b[34;49;3;1m
MAGENTA = \x1b[35;49;3;1m
CYAN = \x1b[36;49;3;1m
WHITE = \x1b[49;3;1m

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

build-dev: swag ## Build the dev binary file
	@echo "${GREEN}build binary file"
	@go build -tags dev -v -a -ldflags="-X 'ggurugi/pkg/cli.GitInfo=$(GIT_INFO)' -X 'ggurugi/pkg/cli.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(DEV)'" -o bin/ggurugi main.go
.PHONY: build-dev

build-prod: swag ## Build the prod binary file
	@echo "${GREEN}build binary file"
	@go build -tags prod -a -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(PROD)'" -o bin/ggurugi main.go
.PHONY: build-prod

run-dev: ### run the dev
	@echo "${YELLOW}run the swag binary file"
	go run -tags dev -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(DEV)'" main.go
.PHONY: run-dev

run-prod: ## run the prod
	@echo "${YELLOW}run the swag binary file"
	go run -tags prod -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'ggurugi/pkg/cli.Config=$(PROD)'" main.go
.PHONY: run-prod
