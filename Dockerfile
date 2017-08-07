FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

ENV GIN_MODE release
ENV GOOGLE_APPLICATION_CREDENTIALS ./credential.json
CMD ["go-wrapper", "run"]
