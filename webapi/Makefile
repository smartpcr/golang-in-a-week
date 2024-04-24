include .env
export $(shell sed 's/=.*//' .env)

echo:
	echo ${MONGO_DB_URI}

build: build-server build-client

build-server:
	CGO_ENABLED=0 go build -o ./bin/server -ldflags '-extldflags "-static"' ./cmd/api/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/tasks-server:dev -f ./cmd/server/Dockerfile .

build-client:
	CGO_ENABLED=0 go build -o ./bin/client -ldflags '-extldflags "-static"' ./cmd/cli/main.go
	cp ${GOBIN}/dlv bin/dlv
	docker build -t smartpcr/tasks-client:dev -f ./cmd/client/Dockerfile .

start-postgres:
	docker run --name tasks-postgres \
	-e POSTGRES_USER=${PG_USER} \
	-e POSTGRES_PASSWORD=${PG_PASSWORD} \
	-e POSTGRES_DB=${PG_DB_NAME} \
	-p ${PG_DB_PORT}:5432 \
	-d postgres:13

stop-postgres:
	docker stop tasks-postgres
	docker rm tasks-postgres

start-mongo:
	docker run --name tasks-mongo -p ${MONGO_DB_PORT}:27017 -d mongo:4.4 \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_DB_USER} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_DB_PASSWORD}

stop-mongo:
	docker stop tasks-mongo
	docker rm tasks-mongo

start-mysql:
	docker run --name tasks-mysql \
	-e MYSQL_ROOT_PASSWORD=${MYSQL_DB_PASSWORD} \
	-e MYSQL_USER=${MYSQL_DB_USER} \
	-e MYSQL_DATABASE=${MYSQL_DB_NAME} \
	-p ${MYSQL_DB_PORT}:3306 \
	-d mysql:8

stop-mysql:
	docker stop tasks-mysql
	docker rm tasks-mysql