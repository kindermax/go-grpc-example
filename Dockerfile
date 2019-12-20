FROM golang:1.13.5-alpine3.10

ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src

COPY . .

# install all go.mod deps
RUN go get ./...
