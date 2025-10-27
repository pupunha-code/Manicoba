FROM golang:1.23.5-alpine AS builder
LABEL authors="kayky"

WORKDIR /app

COPY go.mod go.sum ./

RUN ["go", "mod", "tidy"]

COPY . .

RUN ["go", "build", "-o", "manicoba"]

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/manicoba .

CMD ["./manicoba"]