# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build_base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /ova-hobby-api ./cmd/ova-hobby-api/main.go


FROM alpine:3.14

RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=build_base /ova-hobby-api ./
CMD ["./ova-hobby-api"]
