FROM golang:1.8

RUN go get github.com/itsubaki/gostream

ENTRYPOINT gostream

EXPOSE 1234
