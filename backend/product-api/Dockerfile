FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -v -d ./...

CMD ["go","run","main.go"]