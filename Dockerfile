FROM golang:1.13-alpine

RUN apk add --no-cache git
RUN go get github.com/beego/bee

CMD ["go", "version"]