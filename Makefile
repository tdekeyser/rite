base_version := $(shell cat version.txt)
commit_count = $(shell git rev-list HEAD --count | tr -d '\n')

version = $(base_version).$(commit_count)

build:
	go build -ldflags "-X main.version=$(version)" -o rite main.go

test:	build
	export RITE_APP_DIR=adapter/webapp/
	go test ./...

clean:
	rm -f rite
