FROM golang:1.6

RUN mkdir /app

ADD . /go/src/github.com/turnercode/zoltron-fe
WORKDIR /go/src/github.com/turnercode/zoltron-fe

RUN go get .
RUN go install .

ADD public /go/bin/public

ENV PORT 3000
ENTRYPOINT /go/bin/zoltron-fe
EXPOSE 3000