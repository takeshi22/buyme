FROM golang:1.15.6
RUN mkdir -p /go/src/timememo

WORKDIR /go/src/timememo/server
ADD . /go/src/timememo
ENV GO156MODULE=on

# CMD ["go", "run", "main.go"]