GO_GET= go get
GH_MODS= ./v3

default: test

testdeps:
	$(GO_GET) github.com/onsi/ginkgo
	$(GO_GET) github.com/onsi/gomega

test:
	@go test -v=1 $(GH_MODS)

vet:
	@go tool vet -all $(GH_MODS)

benchmark:
	@go test -test.bench="." -test.run="-" $(GH_MODS)
