FROM golang:latest

WORKDIR /app

COPY . /app

EXPOSE 8080

RUN go mod init main.go
RUN go mod tidy
RUN go run .