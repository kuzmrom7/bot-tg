FROM golang:1.9

MAINTAINER kuzmrom7

RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/github.com/kuzmrom7/bot-tg/src/cmd
ADD . /go/src/github.com/kuzmrom7/bot-tg

RUN govendor sync

RUN go build

ENTRYPOINT ["/go/src/github.com/kuzmrom7/bot-tg/src/cmd/cmd"]