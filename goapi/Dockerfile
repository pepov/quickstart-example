FROM golang:1.20-alpine

ENV PORT=8080
EXPOSE ${PORT}
WORKDIR /app

RUN apk add entr

COPY main.go .