FROM library/golang:1.7.4-alpine

RUN apk update && apk add git
RUN go get github.com/alanchchen/golang-example && \
    go install github.com/alanchchen/golang-example

CMD ["golang-example"]
