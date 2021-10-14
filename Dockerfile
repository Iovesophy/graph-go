FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=auto \
    && cd /app && go test -v
