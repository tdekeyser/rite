# rite

![Go](https://github.com/tdekeyser/rite/workflows/Go/badge.svg?branch=master)

Run the package locally in Docker:

	docker run -p 8080:8080 golang sh -c \
		"go get github.com/tdekeyser/rite/... && export RITE_APP_DIR=/go/src/github.com/tdekeyser/rite/adapter/webapp/ && exec rite"
