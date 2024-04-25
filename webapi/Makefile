.DEFAULT_GOAL:=help

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


## --------------------------------------
## Environment
## --------------------------------------
##@ Environment:
include .env
env: ## load .env variables
	@pwsh -Command '$$content = Get-Content ".\.env"; $$content | ForEach-Object {$$split = $$_.Split("="); if ($$split.Length -gt 1) { [System.Environment]::SetEnvironmentVariable($$split[0], $$split[1]) } }'
	@echo "Environment variables loaded"
	@echo "Db type: ${DB_TYPE}"


## --------------------------------------
## Generator
## --------------------------------------
##@ Generator:

.PHONY: install-wire
install-wire: ## install wire
ifeq (${OS},Linux)
	@which wire > /dev/null || go install github.com/google/wire/cmd/wire@latest
else ifeq (${OS},Windows_NT)
	@where wire >nul 2>nul || go install github.com/google/wire/cmd/wire@latest
else
	$(error Unknown OS: ${OS})
endif

.PHONE: generate
wire-gen: env install-wire ## generate wire
	cd ./pkg/inject && wire

## --------------------------------------
## Build
## --------------------------------------
##@ build:

.PHONY: build
build: build-server build-client ## build server and client

.PHONY: build-server
build-server: wire-gen ## build server based on ${OS}
ifeq (${OS},Linux)
	@$(MAKE) build-server-linux
else ifeq (${OS},Windows_NT)
	@$(MAKE) build-server-windows
else
	$(error Unknown OS: ${OS})
endif

.PHONY: build-server-linux
build-server-linux: ## build server on linux
	CGO_ENABLED=0 go build -o ./bin/server -ldflags '-extldflags "-static"' ./cmd/api/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/tasks-server:dev -f ./cmd/server/Dockerfile .

.PHONY: build-server-windows
build-server-windows: ## build server on windows
	go build -o ./bin/server.exe ./cmd/api/main.go

.PHONY: build-client
build-client: env ## build client based on ${OS}
ifeq (${OS},Linux)
	@$(MAKE) build-client-linux
else ifeq (${OS},Windows_NT)
	@$(MAKE) build-client-windows
else
	$(error Unknown OS: ${OS})
endif

.PHONY: build-client-linux
build-client-linux: env ## build client on linux
	CGO_ENABLED=0 go build -o ./bin/client -ldflags '-extldflags "-static"' ./cmd/cli/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/tasks-client:dev -f ./cmd/client/Dockerfile .

.PHONY: build-client-windows
build-client-windows: env ## build client on windows
	go build -o ./bin/client.exe ./cmd/cli/main.go

.PHONY: run-server
run-server: env build-server start-db ## start server
	./bin/server

## --------------------------------------
## DB
## --------------------------------------
##@ Database

start-db: env ## start database based on ${DB_TYPE}
ifeq (${DB_TYPE},postgres)
	@$(MAKE) start-postgres
else ifeq (${DB_TYPE},mongo)
	@$(MAKE) start-mongo
else ifeq (${DB_TYPE},mysql)
	@$(MAKE) start-mysql
else
	$(error Unknown DB_TYPE: ${DB_TYPE})
endif

stop-db: env ## stop database based on ${DB_TYPE}
ifeq (${DB_TYPE},postgres)
	@$(MAKE) stop-postgres
else ifeq (${DB_TYPE},mongo)
	@$(MAKE) stop-mongo
else ifeq (${DB_TYPE},mysql)
	@$(MAKE) stop-mysql
else
	$(error Unknown DB_TYPE: ${DB_TYPE})
endif

.PHONY: start-postgres
start-postgres: env ## start postgres database
	docker run --name tasks-postgres \
    	-e POSTGRES_USER=${PG_DB_USER} \
    	-e POSTGRES_PASSWORD=${PG_DB_PASSWORD} \
    	-e POSTGRES_DB=${PG_DB_NAME} \
    	-p ${PG_DB_PORT}:5432 \
    	-d postgres:13

.PHONY: stop-postgres
stop-postgres: env ## stop postgres database
	docker stop tasks-postgres
	docker rm tasks-postgres

.PHONY: start-mongo
start-mongo: env ## start mongo database
	docker run --name tasks-mongo -p ${MONGO_DB_PORT}:27017 -d mongo:4.4 \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_DB_USER} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_DB_PASSWORD}

.PHONY: stop-mongo
stop-mongo: env ## stop mongo database
	docker stop tasks-mongo
	docker rm tasks-mongo

.PHONY: start-mysql
start-mysql: env ## start mysql database
	docker run --name tasks-mysql \
	-e MYSQL_ROOT_PASSWORD=${MYSQL_DB_PASSWORD} \
	-e MYSQL_USER=${MYSQL_DB_USER} \
	-e MYSQL_DATABASE=${MYSQL_DB_NAME} \
	-p ${MYSQL_DB_PORT}:3306 \
	-d mysql:8

.PHONY: stop-mysql
stop-mysql: env ## stop mysql database
	docker stop tasks-mysql
	docker rm tasks-mysql


## --------------------------------------
## run
## --------------------------------------
##@ run

.PHONY: run
run: ## run server and client
ifeq (${OS},Linux)
	@$(MAKE) run-server-linux
else ifeq (${OS},Windows_NT)
	@$(MAKE) run-server-windows
else
	$(error Unknown OS: ${OS})
endif

.PHONY: run-server-windows
run-server-windows: env build-server-windows start-db ## run server on windows
	./bin/server.exe

.PHONY: run-server-linux
run-server-linux: env build-server-linux start-db ## run server on linux
	./bin/server

.PHONY: run-client-windows
run-client-windows: env build-client-windows ## run client on windows
	./bin/client.exe

.PHONY: run-client-linux
run-client-linux: env build-client-linux ## run client on linux
	./bin/client

## --------------------------------------
## management dashboard
## --------------------------------------
##@ dashboard

.PHONY: start-portainer-windows
start-portainer-windows: ## start portainer on windows
	docker volume create portainer_data
	docker run -d \
		-p 8000:8000 \
		-p 9443:9443 \
		--name portainer \
		--restart always \
		-v //./pipe/docker_engine://./pipe/docker_engine \
		-v portainer_data:E:\\docker\\portainer_data \
		portainer/portainer-ee:latest

.PHONY: stop-portainer-windows
stop-portainer-windows: ## stop portainer on windows
	docker stop portainer
	docker rm portainer
	docker volume rm portainer_data
