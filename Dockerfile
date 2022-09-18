FROM golang:1.19-alpine

WORKDIR /bot

COPY . /bot

RUN go mod init main.go
RUN go mod tidy