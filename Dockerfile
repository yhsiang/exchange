FROM golang:latest

WORKDIR $GOPATH/src/github.com/yhsiang/exchange
COPY . $GOPATH/src/github.com/yhsiang/exchange
RUN make tools
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./exchange"]