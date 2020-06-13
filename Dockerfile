FROM golang:latest

WORKDIR $GOPATH/src/sec-noti
COPY . $GOPATH/src/sec-noti

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -v

EXPOSE 7000
ENTRYPOINT ["./sec-noti"]
