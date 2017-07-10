FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/gigmsn/publisher/

ADD main.go .

ADD server/ server/

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

ENTRYPOINT ["publisher"]
