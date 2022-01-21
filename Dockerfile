FROM golang:alpine

ADD models.go /usr/local/go/src/form3/account/models/models.go
ADD src /usr/local/go/src/form3/account/client

WORKDIR /go/src/main

ADD main.go .

RUN go mod init

RUN go install -v ./...

CMD ["main"]
