FROM golang:1.12

WORKDIR /app

ADD . /app

RUN go build -mod vendor

CMD ./run.sh

