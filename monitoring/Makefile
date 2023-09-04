BUF_REGISTRY ?= buf.build
BUF_REGISTRY_USERNAME ?= smartpcr
SERVICE_NAME ?= monitoring

setup:
	@./init.sh

publish-schema: setup
	include ./proto/.env.make
	cd ./proto
	buf mod init ${BUF_REGISTRY}/${BUF_REGISTRY_USERNAME}/${SERVICE_NAME}
	echo $BUF_BUILD_TOKEN | buf registry login ${BUF_REGISTRY} --username ${BUF_REGISTRY_USERNAME} --token-stdin
	buf build
	buf push