BUF_REGISTRY ?= buf.build
BUF_REGISTRY_USERNAME ?= smartpcr
SERVICE_NAME ?= monitoring
BUF_TAG ?= $(shell date ++%Y%m%d%H%M)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.DEFAULT_GOAL:=help

## --------------------------------------
## Help
## --------------------------------------

##@ Help:

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


## --------------------------------------
## Generate
## --------------------------------------
##@ Generate:

.PHONY: generate-stub
generate-stub: ## generate stub for go and c#
	cd ./proto
	buf mod init ${BUF_REGISTRY}/${BUF_REGISTRY_USERNAME}/${SERVICE_NAME}
	rm -rf ./gen
	buf build
	buf generate


## --------------------------------------
## Setup
## --------------------------------------
##@ Setup:

.PHONY: pre-requirement
pre-requirement: ## install protoc-gen-go and buf
	@./pre-req.sh

.PHONY: setup
setup: ## retrieve token for buf.build
	@./init.sh

.PHONY: publish-proto
publish-proto: setup ## publish proto contract to buf.build
	cd ./proto
	include .env.make
	cd ./proto
	buf mod init ${BUF_REGISTRY}/${BUF_REGISTRY_USERNAME}/${SERVICE_NAME}
	rm -rf ./gen
	buf build
	buf generate
	echo $BUF_BUILD_TOKEN | buf registry login ${BUF_REGISTRY} --username ${BUF_REGISTRY_USERNAME} --token-stdin
	buf push --tag ${BUF_TAG}


## --------------------------------------
## Certs
## --------------------------------------
##@ Certs:

.PHONY: generate-certs
generate-certs: ## Generate certificate for service authentication
	@./certs.sh


## --------------------------------------
## Build
## --------------------------------------
##@ Build:

.PHONY: build-server
build-server: ## Build server
	CGO_ENABLED=0 go build -o ./bin/server -ldflags '-extldflags "-static"' ./server/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/monitoring-server:dev -f ./server/Dockerfile .

.PHONY: build-client
build-client: ## Build client
	CGO_ENABLED=0 go build -o ./bin/client -ldflags '-extldflags "-static"' ./client/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/monitoring-client:dev -f ./client/Dockerfile .

## --------------------------------------
## Run
## --------------------------------------
##@ Run:

.PHONY: run-local
run-local: start-monitoring-stack ## Run server and client
	go run ./server/main.go
	go run ./client/main.go

.PHONY: run-docker
run-docker: generate-certs build-server build-client ## Run monitoring stack: grafana, prometheus, open-telemetry-collector, jaeger-all-in-one
	docker-compose up -d
