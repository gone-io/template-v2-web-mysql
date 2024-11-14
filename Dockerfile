FROM golang:1.22-alpine3.19 as builder


WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN sed 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' -i /etc/apk/repositories && \
    apk add git && \
    go install github.com/gone-io/gonectr@latest && \
    go install go.uber.org/mock/mockgen@latest

COPY ["go.mod", "go.sum", "Makefile", "./"]

RUN --mount=type=cache,target=/go,id=go go mod download

COPY cmd cmd
COPY internal internal

#编译
RUN --mount=type=cache,target=/go,id=go  go generate ./... && \
    go build -ldflags="-w -s" -tags musl -o bin/server ./cmd/server

FROM alpine:3.19
RUN sed 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' -i /etc/apk/repositories && \
    apk update && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
COPY config config
COPY --from=builder /app/bin/server /app/server

ARG ENVIRONMENT=dev
ENV ENV=${ENVIRONMENT}

CMD ./server
EXPOSE 8080