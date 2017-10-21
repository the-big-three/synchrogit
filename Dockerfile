FROM golang:alpine

RUN apk update && apk add git

ADD . /go/src/github.com/the-big-three/synchrogit

WORKDIR /go/src/github.com/the-big-three/synchrogit

RUN go install

CMD ["synchrogit"]