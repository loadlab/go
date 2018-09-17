FROM golang:1.8

WORKDIR /go/src/loadlab

ADD ./main.go /go/src/loadlab

RUN go get github.com/gorilla/http/client

RUN go build /go/src/loadlab/main.go