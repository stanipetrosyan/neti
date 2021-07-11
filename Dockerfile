FROM golang:1.16.5-alpine3.13

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go build -o server .

CMD ["/app/server"]
