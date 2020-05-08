FROM golang:latest

ADD . /go/src/app/
WORKDIR /go/src/app/

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure 
RUN go install app

ENTRYPOINT /go/bin/app

EXPOSE 4000
