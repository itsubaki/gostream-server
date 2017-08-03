FROM golang:1.8

RUN go get github.com/itsubaki/gostream

ENV GOSTREAM_PORT 80
ENTRYPOINT gostream

EXPOSE 80
