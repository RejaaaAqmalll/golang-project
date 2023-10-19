FROM golang:1.20.4 AS build-env
WORKDIR /app
RUN apk update && apk add --no-cache git
COPY . .
RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]