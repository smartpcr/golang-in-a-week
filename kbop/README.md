mkdir kbop && cd kbop

go mod init tutorials.smartpcr.io/kbop

kubebuilder init --domain tutorials.smartpcr.io --repo tutorials.smartpcr.io/kbop

kubebuilder create api --group core --version v1 --kind Pod