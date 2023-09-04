#! /bin/bash

export TENANT_ID="625a8c92-2669-4d71-8ac3-923a55242192"
export SUBSCRIPTION_ID="c5a015e6-a59b-45bd-a621-82f447f46034"
export RESOURCE_GROUP="xd"
export BUF_REGISTRY="buf.build"
export SERVICE_NAME="monitoring"
export BUF_REGISTRY_USERNAME="smartpcr"
export VAULT_NAME="akshci-kv-xiaodong"
export BUF_BUILD_TOKEN_SECRET_NAME="buf-builder-token"

az login --tenant $TENANT_ID --use-device-code --allow-no-subscriptions
az account set --subscription $SUBSCRIPTION_ID
export BUF_BUILD_TOKEN=$(az keyvault secret show --vault-name $VAULT_NAME --name $BUF_BUILD_TOKEN_SECRET_NAME --query value -o tsv)
echo "BUF_BUILD_TOKEN=${BUF_BUILD_TOKEN}" > ./proto/.env.make

pushd .
cd ./proto
buf mod init ${BUF_REGISTRY}/${BUF_REGISTRY_USERNAME}/${SERVICE_NAME}
echo $BUF_BUILD_TOKEN | buf registry login ${BUF_REGISTRY} --username ${BUF_REGISTRY_USERNAME} --token-stdin
buf build
buf push
popd