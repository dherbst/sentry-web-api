.PHONY: all clean pull get build build-in-container test test-in-container sec sec-in-container lint lint-in-container

GOLANG := golang:1.15
GOOS := darwin

VERSION := 1.0.0
GIT_HASH = $(shell git rev-parse --short HEAD)
LDFLAGS := "-X github.com/dherbst/sentry.GitHash=${GIT_HASH} -X github.com/dherbst/sentry.Version=${VERSION}"

all: clean pull lint sec test build install

clean:
	mkdir -p bin
	rm -f bin/sentry || true

pull:
	docker pull ${GOLANG}

lint:
	docker run -i --rm -v ${PWD}:/go/src/github.com/dherbst/sentry-web-api -w /go/src/github.com/dherbst/sentry-web-api ${GOLANG} make lint-in-container

lint-in-container:
	go get -u golang.org/x/lint/golint
	golint github.com/dherbst/sentry
	golint github.com/dherbst/sentry-web-api/cmd/sentry/...

sec:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/sentry-web-api -w /go/src/github.com/dherbst/sentry-web-api ${GOLANG} make sec-in-container

sec-in-container:
	go get -u github.com/securego/gosec/cmd/gosec
	gosec .

test:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/sentry-web-api -w /go/src/github.com/dherbst/sentry-web-api ${GOLANG} make test-in-container

test-in-container:
	go test -ldflags ${LDFLAGS} -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html

build:
	docker run -i --rm -v "$(PWD)":/go/src/github.com/dherbst/sentry-web-api -w /go/src/github.com/dherbst/sentry-web-api ${GOLANG} make build-in-container

build-in-container:
	GOOS=darwin go build -o bin/sentry -ldflags ${LDFLAGS} cmd/sentry/*.go

install:
	mkdir -p ~/bin
	cp bin/sentry ~/bin/sentry

install-local:
	go install -ldflags ${LDFLAGS} github.com/dherbst/sentry-web-api/cmd/sentry

# gh-release creates a new release in github
gh-release:
	gh release create ${VERSION} 'bin/sentry'
