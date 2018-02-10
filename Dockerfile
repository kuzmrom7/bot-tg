FROM golang:1.9

MAINTAINER kuzmrom7

RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/github.com/kuzmrom7/proxy/cmd/proxy
ADD . /go/src/github.com/lavrs/proxy

RUN govendor sync

RUN go build

ENTRYPOINT ["/go/bin/src/cmd"]