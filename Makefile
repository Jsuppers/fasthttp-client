PROJECT_NAME := "fasthttp-client"

.PHONY: build lint

all: build lint coverage

dep:
	go get -v -d ./...

generate: dep
	@go generate ./...

build: generate
	CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o ./bin/${PROJECT_NAME}

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
	bin/golangci-lint run

coverage: generate
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out
	goveralls -coverprofile=cover.out -service=travis-ci -repotoken ${COVERALLS_TOKEN}
	rm cover.out

coverhtml: generate
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html
	rm cover.out
