FROM golang:latest

ADD . /go/src/app/

RUN go get labix.org/v2/mgo
RUN go get labix.org/v2/mgo/bson
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/go-macaron/binding
RUN go get github.com/go-macaron/session
RUN go get gopkg.in/macaron.v1

RUN go install app

ENTRYPOINT /go/bin/app

EXPOSE 4000
