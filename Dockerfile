FROM golang:alpine

RUN mkdir /app
ADD . /app/
RUN apk update && apk add git && go get github.com/go-martini/martini

WORKDIR /app
RUN go build -o main .

EXPOSE 3000

ENTRYPOINT /app/main
