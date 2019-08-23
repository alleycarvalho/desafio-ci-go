FROM golang:1.12-alpine

WORKDIR /go/src/calculadora

ADD ./src/calculadora /go/src/calculadora

RUN go install .

CMD ["calculadora"]
