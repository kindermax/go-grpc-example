FROM golang:1.13.5-alpine3.10

ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/app

COPY . .

# install all go.mod deps
RUN go get ./...

# install lint tools
RUN go get -u golang.org/x/lint/golint
# for tests https://github.com/golang/go/issues/26988
RUN go get -u golang.org/x/net
