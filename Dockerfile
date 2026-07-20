# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/api ./cmd/api

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /bin/api /app/api

ENV APP_NAME=go-hexagonal-template
ENV APP_ENV=production
ENV HTTP_PORT=8080
ENV CORS_ALLOWED_ORIGINS=*

EXPOSE 8080

ENTRYPOINT ["/app/api"]
