FROM golang:alpine AS builder
COPY . goApp/.
WORKDIR goApp
RUN go mod download
RUN go build main.go
RUN ls -l
ENTRyPOINT ["/go/goApp/main"]