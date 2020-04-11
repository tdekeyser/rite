commit_count = $(shell git rev-list HEAD --count | tr -d '\n')

build:
	go build -ldflags "-X main.commits=.$(commit_count)" -o $(GOROOT)/bin/rite main.go

test:	build
	export RITE_APP_DIR=adapter/webapp/
	go test ./...

clean:
	rm -f $(GOROOT)/bin/rite
