BUF_VERSION=v1.26.1
TAG?=dev


.DEFAULT_GOAL:=help

## --------------------------------------
## Help
## --------------------------------------

##@ Help:

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

guard-%:
	@ if [ "${${*}}" = "" ]; then \
        echo "Environment variable $* not set"; \
        exit 1; \
    fi

install-buf: ## install buf
	go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}

lint: guard-GOPATH ## lint
	${GOPATH}/bin/buf lint

gen-go: install-buf guard-GOPATH ## generate go types
	${GOPATH}/bin/buf generate --template buf.gen.go.yaml

gen-go-server: install-buf guard-GOPATH ## generate go grpc code
	${GOPATH}/bin/buf generate --template buf.gen.go-server.yaml

push-tag: ## push tag to buf
	cd protobuf && buf push --tag=${TAG}