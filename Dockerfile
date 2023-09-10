FROM golang:latest

RUN mkdir -p /go/src/github.com/alizoubair/go-grpc-server

WORKDIR /go/src/github.com/alizoubair/go-grpc-server

COPY . /go/src/github.com/alizoubair/go-grpc-server

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-grpc-server