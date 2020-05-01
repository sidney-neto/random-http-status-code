FROM golang:alpine

RUN mkdir -p /go/src/random-http-status-code
WORKDIR /go/src/random-http-status-code

COPY . /go/src/random-http-status-code

RUN apk add --no-cache git

RUN go get -u github.com/gorilla/mux

CMD ["go", "run", "webserver.go"]
