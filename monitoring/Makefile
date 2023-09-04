BUF_REGISTRY ?= buf.build
BUF_REGISTRY_USERNAME ?= smartpcr
SERVICE_NAME ?= monitoring
BUF_TAG ?= $(shell date ++%Y%m%d%H%M)

setup:
	@./init.sh

publish-schema: setup
	include ./proto/.env.make
	cd ./proto
	buf mod init ${BUF_REGISTRY}/${BUF_REGISTRY_USERNAME}/${SERVICE_NAME}
	echo $BUF_BUILD_TOKEN | buf registry login ${BUF_REGISTRY} --username ${BUF_REGISTRY_USERNAME} --token-stdin
	buf build
	buf push --tag ${BUF_TAG}