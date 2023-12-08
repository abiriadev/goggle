# syntax=docker/dockerfile:1

FROM golang:1.21.5-alpine3.18 as indexer

WORKDIR /app

COPY go.mod go.sum ./

RUN ["go", "mod", "download"]

COPY ./pkg ./pkg
COPY ./cmd ./cmd

RUN ["go", "run", "./cmd/indexer"]

RUN ["go", "build", "./cmd/goggle"]

FROM alpine:3.18 as runner

WORKDIR /app

USER nonroot:nonroot

COPY --from=indexer /app/index.gob /app/goggle ./

EXPOSE 6099

CMD ["./goggle"]
