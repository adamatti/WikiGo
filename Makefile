.DEFAULT_GOAL := help

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## remove binary and coverage files
	@rm -f wiki
	@rm -f *coverage.html

build: ## make binary
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo

compile: build

run: build ## build and run
	@./wiki

compress: build ## compress binary
	@upx wiki

test: clean ## run tests
	@go test -v ./... -coverprofile coverage.html

show-test-coverage: ## show test coverage
	@go tool cover -html=coverage.html

docker-build: ## build docker image
	@docker build -t wiki-go:latest .

docker-build-builder:
	@docker build --target builder -t wiki-go-builder:latest .

docker-run: docker-build ## run docker image
	@docker run --rm -p 3000:3000 wiki-go:latest

docker-ssh: docker-build-builder ## enter bash inside docker containter
	@docker run -it --rm wiki-go-builder:latest /bin/sh