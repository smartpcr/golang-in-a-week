mkdir -p ~/work/github/go/golang-in-a-week/kbop && cd ~/work/github/go/golang-in-a-week/kbop
go mod init tutorial.smartpcr.io/project
kubebuilder init --domain tutorial.smartpcr.io --repo tutorial.smartpcr.io/project

kubebuilder create api --group batch --version v1 --kind Pod