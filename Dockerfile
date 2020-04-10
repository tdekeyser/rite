FROM golang:1.13

COPY . /go/src/github.com/tdekeyser/rite/

WORKDIR /go/src/github.com/tdekeyser/rite/

RUN make build

ENV RITE_APP_DIR=/go/src/github.com/tdekeyser/rite/adapter/webapp/

EXPOSE 8080

CMD ["rite"]
