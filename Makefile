default: vet test

test:
	go test -v ./...

vet:
	go vet ./...

bench:
	go test -bench=. -benchmem -test.run=NONE ./...
