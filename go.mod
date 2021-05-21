module github.com/cloud-team-poc/cluster-patch

go 1.16

require (
	github.com/metal3-io/cluster-api-provider-metal3 v0.4.1
	github.com/onsi/ginkgo v1.14.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200930160638-afb6bcd081ae // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	k8s.io/apimachinery v0.19.0
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/controller-runtime v0.6.2
)

replace k8s.io/client-go => k8s.io/client-go v0.19.0
