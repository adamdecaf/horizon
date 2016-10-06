.PHONY: vet build test

vet:
	./scripts/cmd.sh $(proj) go tool vet .

build: vet
	./scripts/cmd.sh $(proj) go build .

test: build
	./scripts/cmd.sh $(proj) go test -v ./...
