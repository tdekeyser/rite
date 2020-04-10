base_version := $(shell cat version.txt)
commit_count = $(shell git rev-list HEAD --count | tr -d '\n')

version = $(base_version).$(commit_count)

build:
	go build -ldflags "-X main.version=$(version)" -o $(GOROOT)/bin/rite main.go

test:	build
	export RITE_APP_DIR=adapter/webapp/
	go test ./...

clean:
	rm -f $(GOROOT)/bin/rite

dockerrun:
	docker run -p 8080:8080 --rm golang sh -c \
		"go get github.com/tdekeyser/rite/... && export RITE_APP_DIR=/go/src/github.com/tdekeyser/rite/adapter/webapp/ && exec rite"
