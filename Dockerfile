FROM golang:1.18.7-alpine3.16

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go build -o server .

CMD ["/app/server"]
