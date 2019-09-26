FROM golang:1.13-alpine

RUN apk add --no-cache git
RUN apk add --no-cache make
RUN go get github.com/beego/bee

WORKDIR /go/src/github.com/ever-cifuentes-globant/weather-api

CMD ["go", "version"]