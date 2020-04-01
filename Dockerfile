FROM golang:latest

WORKDIR $GOPATH/src/sec-noti
COPY . $GOPATH/src/sec-noti

RUN go build server.go

EXPOSE 7000
ENTRYPOINT ["./server"]
