default: vet test

testdeps:
	go get -v github.com/onsi/ginkgo
	go get -v github.com/onsi/gomega
	
test:
	go test -v ./...

vet:
	go vet ./...

bench:
	go test -bench=. -benchmem -test.run=NONE ./...
