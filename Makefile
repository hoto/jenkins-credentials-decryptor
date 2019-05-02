all: clean build test

clean:
	go clean
	rm -rf bin/ dist/

dependencies:
	go get -v -t -d ./...
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build: dependencies
	go build -o bin/jenkins-credentials-decryptor cmd/jenkins-credentials-decryptor/main.go

test:
	go test -v ./...

install: clean build
	go install -v ./...

github-release: dependencies
	curl -sL https://git.io/goreleaser | bash

github-release-dry-run: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist
