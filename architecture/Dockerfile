FROM golang:1.18-alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ADD . /app

ARG BUILD_ARG='local'
ENV BUILD_ENV=$BUILD_ARG
ARG GIT_INFO_ARG='default-sha'
ENV GIT_INFO_ENV=$GIT_INFO_ARG
ARG TAG_ARG='v1.0.0'
ENV TAG_ENV=$TAG_ARG


RUN apk add alpine-sdk
RUN go build -v -a -ldflags="-X 'project/pkg/cli.GitInfo=$GIT_INFO_ENV' -X 'project/pkg/cli.Tag=$TAG_ENV' -X 'project/pkg/cli.Config=$BUILD_ENV'" -o bin/project main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/project /app/project
COPY config /app/config

CMD "/app/project"
