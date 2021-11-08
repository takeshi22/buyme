FROM golang:1.15.6
RUN mkdir -p /go/src/buyme

WORKDIR /go/src/buyme/server
ADD . /go/src/buyme
ENV GO156MODULE=on

CMD ["go", "run", "server.go"]