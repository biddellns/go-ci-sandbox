FROM golang:1.11

ADD . /go/src/github.com/biddellns/go-ci-sandbox
WORKDIR /go/src/github.com/biddellns/go-ci-sandbox

CMD ["go", "run", "main.go"]
